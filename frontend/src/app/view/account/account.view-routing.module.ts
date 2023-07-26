import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AccountViewComponent } from './account.view.component';

const routes: Routes = [
    {
        path: '',
        component: AccountViewComponent,
    },
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule],
})
export class AccountViewRoutingModule {}
