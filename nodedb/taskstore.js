"use strict";

// Talking to a MongoDB from NodeJS

const mongodb = require("mongodb");

class MongoStore {
    constructor(db, colName) {
        this.collection = db.collection(colName);
    }
    insert(task) {
        task._id = new mongodb.ObjectID();
        return this.collection.insertOne(task) // this is the promise
            .then(() => task); // but we don't care about the promise, we just want what's actually returned
            // we are not putting a catch here because it's the caller of insert that will handle the error
    }
    update(id, updates) {
        let updateDoc = {
            "$set": updates
        }
        return this.collection.findOneAndUpdate(
            {_id: id}, 
            updateDoc, 
            {returnOriginal: false})
            .then(result => result.value); // result.value is the updated result 
    }
    get(id) {
        return this.collection.findOne({_id: id});
    }
    delete(id) {
        return this.collection.deleteOne({_id: id});
    }
    getAll(completed) { // do we want completed tasks or uncompleted tasks
        return this.collection.find({completed: completed})
            .limit(1000) // limit so we don't get too much
            .toArray(); // will change results to an array to return
    }
}

module.exports = MongoStore;