import { Injectable } from '@angular/core';
import { State, Store } from '@ngrx/store';

import { Paho } from 'ng2-mqtt/mqttws31';
import { WindowRef } from '../service/jarvis-utils.service';
import { LoggerService } from '../service/logger.service';
import { BrokerStoreService, NewMessageAction } from '../store/broker.store';
import { MessageBean } from '../model/broker/message-bean';

@Injectable()
export class JarvisMqttService {
  private _client: Paho.MQTT.Client;

  constructor(
    private brokerStoreService: BrokerStoreService,
    private window: WindowRef,
    private logger: LoggerService
  ) {
    this.logger.warn('MQTT connect', this.window.getHostname(), 8000);
    this._client = new Paho.MQTT.Client(this.window.getHostname(), 8000, "clientId-front");

    this._client.onConnectionLost = (responseObject: Object) => {
      this.logger.error('Connection lost.', responseObject);
      // Reconnect
      this._client.connect({
        onSuccess: this.onConnected.bind(this)
      });
    };

    this._client.onMessageArrived = (message: Paho.MQTT.Message) => {
      this.brokerStoreService.dispatch(new NewMessageAction(
        <MessageBean>{
          id: '',
          topic: message.destinationName,
          body: message.payloadString,
        }
      ));
    };

    this._client.connect({
      onSuccess: this.onConnected.bind(this)
    });
  }

  private onConnected(): void {
    this.logger.info('Connected to broker and subscribe for #');
    this._client.subscribe("#", {});
  }
}
