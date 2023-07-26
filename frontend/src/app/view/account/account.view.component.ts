import { Component } from '@angular/core';
import { RegisterService } from '../../service/register.service';
import { AUTH_STATE } from '../../model/types';
import { SnackbarService } from '../../service/snackbar.service';
import { Observable } from 'rxjs';
import { AppStateService, IAppState } from '../../service/appstate.service';

@Component({
  selector: 'app-account',
  templateUrl: './account.view.component.html',
  styleUrls: ['./account.view.component.scss'],
})
export class AccountViewComponent {
  public appState$: Observable<IAppState>;

  constructor(
    private appStateService: AppStateService,
    private registerService: RegisterService,
    private snackbarService: SnackbarService,
  ) {
    this.appState$ = this.appStateService.getAppState$();
  }

  public registerNewKey(): void {
    this.registerService.registerNewKey().subscribe(authState => {
      authState === AUTH_STATE.VERIFIED ?
        this.snackbarService.showSnackbarWithMessage('Register of new key was successful')
        : this.snackbarService.showSnackbarForAuthState(authState);
    });
  }
}
