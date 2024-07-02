import { Injectable } from '@angular/core';
import { BaseService } from '../baseService';
import { Observable } from 'rxjs';
import { GlobalConstants } from 'src/app/utilities/globalConstants';
import { AddReminder } from 'src/app/interfaces/add-reminder';

@Injectable({
  providedIn: 'root'
})
export class ReminderService {
  private GetRemindersEndpoint : string = GlobalConstants.ApiUrl+"Reminders";
  private AddRemindersEndpoint : string = GlobalConstants.ApiUrl+"Reminders/Add";
  private DeleteRemindersEndpoint : string = GlobalConstants.ApiUrl+"Reminders/Delete";
  private UpdateRemindersEndpoint : string = GlobalConstants.ApiUrl+"Reminders/Update";
  constructor(private baseService : BaseService) { }

  GetReminders(): Observable<any> {
    return this.baseService.makeGetRequest(this.GetRemindersEndpoint);
  }

  AddReminders(reminder: AddReminder): Observable<any>{
    return this.baseService.makePostRequest(this.AddRemindersEndpoint, reminder);
  }
  Updatereminder(reminder: AddReminder): Observable<any>{
    return this.baseService.makePutRequest(this.UpdateRemindersEndpoint, reminder);
  }
  DeleteReminders(reminder: AddReminder): Observable<any>{
    return this.baseService.makeDeleteRequest(this.DeleteRemindersEndpoint+"/"+reminder.reminderId);
  }
}
