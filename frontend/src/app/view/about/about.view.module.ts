import { NgModule } from '@angular/core';
import { AboutViewComponent } from './about.view.component';
import { AboutViewRoutingModule } from './about.view-routing.module';


@NgModule({
  declarations: [
    AboutViewComponent,
  ],
  imports: [
    AboutViewRoutingModule,
  ],
  providers: [],
})
export class AboutViewModule { }
