import { Injectable } from '@angular/core';
import { HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

import { ApiService } from './api.service';

@Injectable({
    providedIn: 'root'
  })
export class PanelService {

    constructor(private apiService: ApiService) { 
    }
  
    get() {
        return this.apiService.get('/logins', 'laravel_be');
    }
  }