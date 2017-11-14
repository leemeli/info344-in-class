"use strict";

const MongoStore = require("./taskstore.js");
const mongodb = require("mongodb");

const mongoAddr = process.env.DBADDR || "localhost:27017"; //have a default set
const mongoURL = `mongodb://${mongoAddr}/tasks`; // put the database name at the end (/tasks)

describe("Mongo Task Store", () => {
    test("CRUD Cycle", () => {
        return mongodb.MongoClient.connect(mongoURL)
        // any time you connect to the database in node, it is asynchronous and you will get a promise
            .then(db => {
                let store = new MongoStore(db, "tasks");
                let task = {
                    title: "Learn Node.js to MongoDB",
                    tags: ["mongodb", "node", "info344"]
                };
                return store.insert(task)
                    .then(task => {
                        expect(task._id).toBeDefined();
                        return task._id;
                    })
                    .then(taskId => {
                        return store.get(taskId); // this gives a promise
                    })
                    .then(fetchedTask => {
                        expect(fetchedTask).toEqual(task);
                        return store.update(task._id, {completed: true});
                    })
                    .then(updatedTask => {
                        expect(updatedTask.completed).toBe(true);
                        return store.delete(task._id);
                    })
                    .then(() => {
                        //testing that the task did properly delete
                        return store.get(task._id);
                    })
                    .then(fetchedTask => {
                        expect(fetchedTask).toBeFalsy(); //either undefined, false, null ,etc
                    })
                    .then(() => {
                        // if all goes well
                        db.close(); 
                    })
                    .catch( err => {
                        db.close();
                        throw err; //If we get here, it will be caught as a test failure
                    });
            });
    });
});