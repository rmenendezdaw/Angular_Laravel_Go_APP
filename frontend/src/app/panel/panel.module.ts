import { ModuleWithProviders, NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { SharedModule } from '../shared';
import { CommonModule } from '@angular/common';
import { PanelComponent } from './panel.component';
import { PanelRoutingModule } from './panel-routing.module';
import { PanelResolver } from './panel-resolver.service';


@NgModule({
  imports: [
    CommonModule,
    SharedModule,
    PanelRoutingModule
  ],
  declarations: [
    PanelComponent
  ],
  providers: [PanelResolver]
})
export class PanelModule { }
