#!/usr/bin/env node
"use strict";

// ^ this lets us chmod how we run bash scripts, program knows how to interpret code
// Node program that will connect to that queue and send messages to the queue

const amqp = require("amqplib");
const qName = "testQ";
const mqAddr = process.env.MQADDR || "localhost:5672";
const mqURL = `amqp://${mqAddr}`;

//anonymous immediately executing function
//async in front of function allows us to call the await
(async function() {
    try {
    console.log("connecting to %s", mqURL);
    let connection = await amqp.connect(mqURL);
    let channel = await connection.createChannel();
    let qConf = await channel.assertQueue(qName, {durable: false});

    console.log("starting to send messages...");
    setInterval(() => {
        let msg = "Message send at " + new Date().toLocaleTimeString();
        channel.sendToQueue(qName, Buffer.from(msg));
        console.log("sending message: Message " +new Date().toLocaleTimeString() );
    }, 1000);
    } catch(err) {
        console.error(err.stack);
    }
})();