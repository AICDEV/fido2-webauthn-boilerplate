import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoggedInGuardCanActivateFn } from './guard/loggedin.guard';


const routes: Routes = [
  {
    path: '*',
    redirectTo: 'login',
    pathMatch: 'full',
  },
  {
    path: 'admin',
    canActivate: [LoggedInGuardCanActivateFn],
    loadChildren: () => import('./view/admin/admin.view.module').then((m) => m.AdminViewModule),
  },
  {
    path: 'main',
    loadChildren: () => import('./view/main/main.view.module').then((m) => m.MainViewModule),
  },
  {
    path: 'about',
    loadChildren: () => import('./view/about/about.view.module').then((m) => m.AboutViewModule),
  },
  {
    path: 'login',
    loadChildren: () => import('./view/login/login.view.module').then((m) => m.LoginViewModule),
  },
  {
    path: 'register',
    loadChildren: () => import('./view/register/register.view.module').then((m) => m.RegisterViewModule),
  },
  {
    canActivate: [LoggedInGuardCanActivateFn],
    path: 'account',
    loadChildren: () => import('./view/account/account.view.module').then((m) => m.AccountViewModule),
  },
  {
    path: '**',
    redirectTo: 'login',
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
