"use strict";

const MongoStore = require("./taskstore.js");
const mongodb = require("mongodb");
const express = require("express");
const app = express(); // create application
const mongoAddr = process.env.DBADDR || "localhost:27017"; //have a default set
const mongoURL = `mongodb://${mongoAddr}/tasks`; // put the database name at the end (/tasks)
const addr = process.env.ADDR || "localhost:4000";
const [host, port] = addr.split(":");


mongodb.MongoClient.connect(mongoURL)
    .then(db => {
        let taskStore = new MongoStore(db, "tasks");

        //parses posted JSON and makes it available from req.body
        app.use(express.json());

        app.post("/v1/tasks", (req, res) => {
            //insert a new task
            taskStore.insert(req.body)
            .then(task => {
                res.json(task);
            })
            .catch(err => {
                throw err;
            });
        });
        
        app.get("/v1/tasks", (req, res) => {
            //return all not-completed tasks in the database
            taskStore.getAll(false)
                .then(tasks => {
                    res.json(tasks);
                })
                .catch(err => {
                    throw err;
                });
        });
        
        app.patch("/v1/tasks/:taskID", (req, res) => {
            let taskIDToFetch = req.params.taskID;
            //update single task by id and send to client
        });
        
        app.listen(port, host, () => {
            console.log(`server is listening at http://${addr}...`);
        });

        db.close();
    });