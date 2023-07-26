import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { AppStateService, IAppState } from '../../service/appstate.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  public appState$: Observable<IAppState>;

  constructor(
    private appStateService: AppStateService,
    private router: Router,
  ) {
    this.appState$ = this.appStateService.getAppState$();
  }

  public logout(): void {
    console.log('logout');
    this.appStateService.logout();
    this.router.navigate(['/main']);
  }
}
