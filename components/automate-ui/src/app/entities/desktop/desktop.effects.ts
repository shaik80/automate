import { catchError, mergeMap, map, withLatestFrom, switchMap, tap } from 'rxjs/operators';
import { Injectable } from '@angular/core';
import { Actions, Effect, ofType } from '@ngrx/effects';
import { of, combineLatest, from } from 'rxjs';
import { NgrxStateAtom } from 'app/ngrx.reducers';
import { Store } from '@ngrx/store';
import { NodeRunsService } from 'app/services/node-details/node-runs.service';
import { NodeRun } from 'app/types/types';

import {
  GetDailyCheckInTimeSeries,
  GetDailyCheckInTimeSeriesSuccess,
  GetDailyCheckInTimeSeriesFailure,
  GetDailyNodeRunsStatusTimeSeries,
  GetDailyNodeRunsStatusTimeSeriesSuccess,
  GetDailyNodeRunsStatusTimeSeriesFailure,
  DesktopActionTypes,
  GetTopErrorsCollection,
  GetTopErrorsCollectionSuccess,
  GetTopErrorsCollectionFailure,
  GetUnknownDesktopDurationCounts,
  GetUnknownDesktopDurationCountsSuccess,
  GetUnknownDesktopDurationCountsFailure,
  GetNodeMetadataCounts,
  GetNodeMetadataCountsSuccess,
  GetNodeMetadataCountsFailure,
  GetDesktops,
  GetDesktopsSuccess,
  GetDesktopsFailure,
  GetDesktop,
  GetDesktopSuccess,
  GetDesktopFailure,
  GetDesktopsTotal,
  GetDesktopsTotalSuccess,
  GetDesktopsTotalFailure,
  UpdateDesktopColumnOptions,
  GetDesktopColumnOptionsDefaults
} from './desktop.actions';
import { DesktopRequests } from './desktop.requests';
import { CreateNotification } from 'app/entities/notifications/notification.actions';
import { Type } from 'app/entities/notifications/notification.model';
import { getSelectedDaysAgo } from 'app/entities/desktop/desktop.selectors';

@Injectable()
export class DesktopEffects {
  constructor(
    private actions$: Actions,
    private requests: DesktopRequests,
    private store$: Store<NgrxStateAtom>,
    private nodeRunsService: NodeRunsService
  ) { }

  @Effect()
  setDaysAgoSelected$ = this.actions$.pipe(
    ofType(DesktopActionTypes.SET_SELECTED_DAYS_AGO),
    map(() => new GetDailyCheckInTimeSeries())
  );

  @Effect()
  getDailyCheckInTimeSeries$ = combineLatest([
    this.actions$.pipe(ofType<GetDailyCheckInTimeSeries>(
      DesktopActionTypes.GET_DAILY_CHECK_IN_TIME_SERIES)),
    this.store$.select(getSelectedDaysAgo)])
    .pipe(
      mergeMap(([_action, selectedDaysAgo]) =>
        this.requests.getDailyCheckInCountCollection(selectedDaysAgo).pipe(
          map(dailyCheckInCountCollection =>
        new GetDailyCheckInTimeSeriesSuccess(dailyCheckInCountCollection)),
        catchError((error) => of(new GetDailyCheckInTimeSeriesFailure(error)))))
    );

  @Effect()
  getDailyCheckInTimeSeriesFailure$ = this.actions$.pipe(
    ofType(DesktopActionTypes.GET_DAILY_CHECK_IN_TIME_SERIES_FAILURE),
    map(({ payload: { error } }: GetDailyCheckInTimeSeriesFailure) => {
      const msg = error.error;
      return new CreateNotification({
        type: Type.error,
        message: `Could not get time series: ${msg || error}`
      });
    }));

  @Effect()
    getDailyNodeRunsStatusTimeSeries$ =
      this.actions$.pipe(ofType<GetDailyNodeRunsStatusTimeSeries>(
        DesktopActionTypes.GET_DAILY_NODE_RUNS_STATUS_TIME_SERIES),
        mergeMap(({ nodeId, daysAgo }) =>
          this.requests.getDailyNodeRunsStatusCountCollection(nodeId, daysAgo).pipe(
            map(dailyNodeRunsStatusCountCollection =>
              new GetDailyNodeRunsStatusTimeSeriesSuccess(dailyNodeRunsStatusCountCollection)),
              catchError((error) => of(new GetDailyNodeRunsStatusTimeSeriesFailure(error))))));

    @Effect()
    getDailyNodeRunsStatusSeriesFailure$ = this.actions$.pipe(
      ofType(DesktopActionTypes.GET_DAILY_NODE_RUNS_STATUS_TIME_SERIES_FAILURE),
      map(({ payload: { error } }: GetDailyNodeRunsStatusTimeSeriesFailure) => {
        const msg = error.error;
        return new CreateNotification({
          type: Type.error,
          message: `Could not get time series: ${msg || error}`
        });
      }));

  @Effect()
  getTopErrorCollection$ =
    this.actions$.pipe(ofType<GetTopErrorsCollection>(
      DesktopActionTypes.GET_TOP_ERRORS_COLLECTION),
      mergeMap((_action) =>
        this.requests.getTopErrorsCollection().pipe(
          map(topErrorsCollection =>
        new GetTopErrorsCollectionSuccess(topErrorsCollection)),
        catchError((error) => of(new GetTopErrorsCollectionFailure(error)))))
    );

  @Effect()
  getTopErrorCollectionFailure$ = this.actions$.pipe(
    ofType(DesktopActionTypes.GET_TOP_ERRORS_COLLECTION_FAILURE),
    map(({ payload: { error } }: GetTopErrorsCollectionFailure) => {
      const msg = error.error;
      return new CreateNotification({
        type: Type.error,
        message: `Could not get top errors: ${msg || error}`
      });
    }));

  @Effect()
  getUnknownDesktopDurationCounts$ =
    this.actions$.pipe(ofType<GetUnknownDesktopDurationCounts>(
      DesktopActionTypes.GET_UNKNOWN_DESKTOP_DURATION_COUNTS),
      mergeMap((_action) =>
        this.requests.getUnknownDesktopDurationCounts().pipe(
          map(countedDurationCollection =>
        new GetUnknownDesktopDurationCountsSuccess(countedDurationCollection)),
        catchError((error) => of(new GetUnknownDesktopDurationCountsFailure(error)))))
    );

  @Effect()
  getNodeMetadataCounts$ = this.actions$.pipe(
    ofType<GetNodeMetadataCounts>(DesktopActionTypes.GET_NODE_METADATA_COUNTS),
    withLatestFrom(this.store$),
    switchMap(([_action, storeState]) =>
      this.requests.getNodeMetadataCounts(storeState.desktops.getDesktopsFilter).pipe(
        map(nodeMetadataCounts => new GetNodeMetadataCountsSuccess(nodeMetadataCounts)),
        catchError((error) => of(new GetNodeMetadataCountsFailure(error))))
    ));

  @Effect()
  getUnknownDesktopDurationCountsFailure$ = this.actions$.pipe(
    ofType(DesktopActionTypes.GET_UNKNOWN_DESKTOP_DURATION_COUNTS_FAILURE),
    map(({ payload: { error } }: GetUnknownDesktopDurationCountsFailure) => {
      const msg = error.error;
      return new CreateNotification({
        type: Type.error,
        message: `Could not get unknown desktop duration counts errors: ${msg || error}`
      });
    }));

  @Effect()
  getDesktops$ = this.actions$.pipe(
    ofType(DesktopActionTypes.GET_DESKTOPS),
    withLatestFrom(this.store$),
    switchMap(([_action, storeState]) =>
      this.requests.getDesktops(storeState.desktops.getDesktopsFilter).pipe(
        map(desktops => new GetDesktopsSuccess( desktops )),
      catchError((error) => of(new GetDesktopsFailure(error))))
    ));

  @Effect()
  getDesktopsFailure$ = this.actions$.pipe(
    ofType(DesktopActionTypes.GET_DESKTOPS_FAILURE),
    map(({ payload: { error } }: GetDesktopsFailure) => {
      const msg = error.error;
      return new CreateNotification({
        type: Type.error,
        message: `Could not get desktops errors: ${msg || error}`
      });
    }));

  @Effect()
  getDesktop$ = this.actions$.pipe(
    ofType(DesktopActionTypes.GET_DESKTOP),
    switchMap((action: GetDesktop) =>
      from(this.nodeRunsService.getNodeRun(action.payload.nodeId, action.payload.runId)).pipe(
        map((nodeRun: NodeRun) => new GetDesktopSuccess(nodeRun)),
        catchError((error) => of(new GetDesktopFailure(error))))
    ));

  @Effect()
  getDesktopFailure$ = this.actions$.pipe(
    ofType(DesktopActionTypes.GET_DESKTOP_FAILURE),
    map(({ payload: { error } }: GetDesktopFailure) => {
      const msg = error.error;
      return new CreateNotification({
        type: Type.error,
        message: `Could not get desktop errors: ${msg || error}`
      });
    }));

  @Effect()
  getDesktopsTotal$ = this.actions$.pipe(
    ofType(DesktopActionTypes.GET_DESKTOPS_TOTAL),
    withLatestFrom(this.store$),
    switchMap(([_action, storeState]) =>
      this.requests.getDesktopsTotal(storeState.desktops.getDesktopsFilter).pipe(
        map(total => new GetDesktopsTotalSuccess( total )),
      catchError((error) => of(new GetDesktopsTotalFailure(error))))
    ));

  @Effect()
  getDesktopsTotalFailure$ = this.actions$.pipe(
    ofType(DesktopActionTypes.GET_DESKTOPS_TOTAL_FAILURE),
    map(({ payload: { error } }: GetDesktopsTotalFailure) => {
      const msg = error.error;
      return new CreateNotification({
        type: Type.error,
        message: `Could not get desktops total errors: ${msg || error}`
      });
    }));

  @Effect()
  updateDesktopFilterCurrentPage$ = this.actions$.pipe(
      ofType(DesktopActionTypes.UPDATE_DESKTOPS_FILTER_CURRENT_PAGE),
      mergeMap(() => [ new GetDesktops() ]));

  @Effect()
  updateDesktopFilterPageSize$ = this.actions$.pipe(
    ofType(DesktopActionTypes.UPDATE_DESKTOPS_FILTER_PAGE_SIZE_AND_CURRENT_PAGE),
    mergeMap(() => [ new GetDesktops() ]));

  @Effect()
  addDesktopFilterTerm$ = this.actions$.pipe(
      ofType(DesktopActionTypes.ADD_DESKTOPS_FILTER_TERM),
      mergeMap(() => [ new GetDesktops(), new GetDesktopsTotal() ]));

  @Effect()
  updateDesktopFilterTerms$ = this.actions$.pipe(
      ofType(DesktopActionTypes.UPDATE_DESKTOPS_FILTER_TERMS),
      mergeMap(() => [ new GetDesktops(), new GetDesktopsTotal() ]));

  @Effect()
  removeDesktopFilterTerms$ = this.actions$.pipe(
      ofType(DesktopActionTypes.REMOVE_DESKTOPS_FILTER_TERM),
      mergeMap(() => [ new GetDesktops(), new GetDesktopsTotal() ]));

  @Effect()
  updateDesktopSortTerm$ = this.actions$.pipe(
      ofType(DesktopActionTypes.UPDATE_DESKTOPS_SORT_TERM),
      mergeMap(() => [ new GetDesktops() ]));

  @Effect({ dispatch: false })
  updateDesktopColumnOptions$ = this.actions$.pipe(
    ofType<UpdateDesktopColumnOptions>(DesktopActionTypes.UPDATE_DESKTOP_COLUMN_OPTIONS),
    tap(({ payload }) => {
      if (payload.saveAsDefault) {
        localStorage.setItem('desktopColumnOptions', JSON.stringify(payload.options));
      }
    }));

  @Effect()
  getDesktopColumnOptionsDefault$ = this.actions$.pipe(
    ofType<GetDesktopColumnOptionsDefaults>(DesktopActionTypes.GET_DESKTOP_COLUMN_OPTIONS_DEFAULTS),
    mergeMap(() => {
      const stored = localStorage.getItem('desktopColumnOptions');
      const options = stored ? JSON.parse(stored) : [];
      return stored ? [new UpdateDesktopColumnOptions({ options, saveAsDefault: true })] : [];
    }));
}
