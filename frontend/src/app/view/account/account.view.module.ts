import { NgModule } from '@angular/core';
import { AccountViewComponent } from './account.view.component';
import { AccountViewRoutingModule } from './account.view-routing.module';
import { MatButtonModule } from '@angular/material/button';
import { CommonModule } from '@angular/common';


@NgModule({
  declarations: [
    AccountViewComponent,
  ],
  imports: [
    CommonModule,
    AccountViewRoutingModule,
    MatButtonModule,
  ],
  providers: [],
})
export class AccountViewModule { }
