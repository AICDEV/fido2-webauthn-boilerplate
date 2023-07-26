import { NgModule } from '@angular/core';
import { AdminViewComponent } from './admin.view.component';
import { AdminViewRoutingModule } from './admin.view-routing.module';
import { MatButtonModule } from '@angular/material/button';
import { NgIf } from '@angular/common';


@NgModule({
  declarations: [
    AdminViewComponent,
  ],
  imports: [
    AdminViewRoutingModule,
    MatButtonModule,
    NgIf,
  ],
  providers: [],
})
export class AdminViewModule { }
