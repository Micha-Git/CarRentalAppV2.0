import {platformBrowserDynamic} from "@angular/platform-browser-dynamic";
import {AppModule} from "./app/app.module";
import {ViewCarInformationRequest,ViewCarInformationResponse,Vin,Car} from "./modules/fleet-management/service/proto/api_specification_am_fleet_management_pb";

platformBrowserDynamic().bootstrapModule(AppModule)
  .catch(reason => {
    // For debug purpose
    console.error(reason);
  });
