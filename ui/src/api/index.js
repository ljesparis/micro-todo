import axios from 'axios';

export function GetTasks(limit, offset, open) {
    return axios.get("/tasks?limit=" + limit + "&offset=" + offset + "&isopen=" + open);
}

export function GetTasksByUrl(url) {
    return axios.get(url);
}

export function GetTask(id) {
    return axios.get("/tasks/" + id);
}

export function CreateTask(name, description, done) {
    return axios.post("/create", {
        name: name,
        description: description,
        done: done,
    });
}

export function UpdateTask(id, name, description, done) {
    return axios.put("/tasks/" + id + "/update", {
        name: name,
        description: description,
        done: done,
    });
}

export function DeleteTask(id) {
    return axios.delete("/tasks/" + id + "/update");
}
