import { CanActivateFn, Router } from '@angular/router';
import { inject } from '@angular/core';
import { AppStateService } from '../service/appstate.service';


export const LoggedInGuardCanActivateFn: CanActivateFn = (): boolean => {
  const router: Router = inject(Router);
  const appStateService: AppStateService = inject(AppStateService);

  if (!appStateService.isLoggedIn()) {
    router.navigate(['login']);
  }
  return appStateService.isLoggedIn();
}
