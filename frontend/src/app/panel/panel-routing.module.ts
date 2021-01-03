import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { PanelComponent } from './panel.component';
import { PanelResolver } from './panel-resolver.service';


const routes: Routes = [
  {
    path: '',
    component: PanelComponent,
    resolve: {
      song: PanelResolver
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PanelRoutingModule {}
