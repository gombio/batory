import { EventEmitter } from 'events';

class Socket {
  constructor(ws = new WebSocket(), ee = new EventEmitter()) {
    this.handlers = [];
    this.ws = ws;
    this.ee = ee;

    ws.onmessage = this.message.bind(this);
    ws.onopen = this.open.bind(this);
    ws.onclose = this.close.bind(this);
  }

  on(name, fn) {
    //Prevent binding one handler multiple times
    let found = false;
    this.handlers.map((handlerName) => {
      if (name === handlerName) {
        found = true;
        return handlerName;
      }
      return handlerName;
    });

    if (found) {
      return;
    }

    this.handlers.push(name);
    this.ee.on(name, fn);
  }

  off(name, fn) {
    this.ee.removeListener(name, fn);
  }

  emit(type, data) {
    const message = JSON.stringify({type, data});
    this.ws.send(message);
  }

  message(e) {
    try {
      const message = JSON.parse(e.data);
      this.ee.emit(message.type, message.data);
    } catch (err) {
      this.ee.emit('error', err);
    }
  }

  open() {
    this.ee.emit('connect');
  }

  close() {
    this.ee.emit('disconnect');
  }
}

export default Socket;
