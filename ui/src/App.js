import React, { Component } from 'react';
import './App.css';

import {
  Container,
  Row,
  Col,
  InputGroup,
  InputGroupAddon,
  Button,
  Input,
} from 'reactstrap';

import {
    GetTasks,
    GetTasksByUrl,
    CreateTask,
    UpdateTask
} from './api';

import Task from './components/task';
import TaskList from './components/tasklist';


class App extends Component {
  constructor(props) {
      super(props);

      this.taskNameInput = React.createRef();
      this.taskDescriptionInput = React.createRef();

      this.state = {
        openedTasks: [],
        closedTasks: [],
        btnActive: false,
        showTaskListCols: true,
        nextOpenedTasksLink: null,
        prevOpenedTasksLink: null,
        nextClosedTasksLink: null,
        prevClosedTasksLink: null,
        currentOpenedTasksUrl: null,
        currentClosedTasksUrl: null,
        isOpenedTasksTabOpen: true,
        limit: 6,
      };

      this._onGetOpenedTasks = this._onGetOpenedTasks.bind(this);
      this._onOpenTasksHandler = this._onOpenTasksHandler.bind(this);
      this._onClosedTasksHandler= this._onClosedTasksHandler.bind(this);
      this._onGetClosedTasks = this._onGetClosedTasks.bind(this);
      this._onErrorHandler = this._onErrorHandler.bind(this);

      this._onResize = this._onResize.bind(this);
      this._onCreateTask = this._onCreateTask.bind(this);
      this._onChangeText = this._onChangeText.bind(this);
      this._onResize = this._onResize.bind(this);
      this._onNextOrPrevt = this._onNextOrPrevt.bind(this);
      this._onTabChanged = this._onTabChanged.bind(this);
  }

  _onTabChanged(t) {
      this.setState({
          ...this.state,
          isOpenedTasksTabOpen: t === "1",
      });
  }

  _onNextOrPrevt(url) {
      let tmp = {};
      if(this.state.isOpenedTasksTabOpen) {
          tmp = {
              currentOpenedTasksUrl: url,
          };
      } else {
          tmp = { currentClosedTasksUrl: url };
      }

      this.setState({
          ...this.state,
          ...tmp,
      });

      GetTasksByUrl(url)
          .then(this.state.isOpenedTasksTabOpen? this._onOpenTasksHandler: this._onClosedTasksHandler)
          .catch(this._onErrorHandler)
  }

  _onCreateTask() {
      let taskName = this.taskNameInput.current.value;
      let taskDescription = this.taskDescriptionInput.current.value;

      if(taskName.length > 0 && taskDescription.length > 0) {
          CreateTask(taskName, taskDescription, false)
              .then(res => {
                  if(res.status === 204) {
                      this._onGetOpenedTasks();
                  }
              })
              .catch(err => {
                  alert("internal server error");
              });

          this.taskNameInput.current.value = "";
          this.taskDescriptionInput.current.value = "";
      }
  }

  _onDone(task) {
      UpdateTask(task.id, task.name, task.description, true)
          .then(res => {
             if(res.status === 200) {
                 this._onGetOpenedTasks();
                 this._onGetClosedTasks();
             }
          })
          .catch(err => {
              alert("internal server error");
          });
  }

  _onChangeText(e) {
      let active = false;
      let des = this.taskDescriptionInput.current.value;
      let n = this.taskNameInput.current.value;
      if(n.length > 0 && des.length > 0 && des.length <= 62) {
          active = true;
      }

      this.setState({
          ...this.state,
          btnActive: active,
      });
  }

  _onResize() {
      let t = true;
      if(window.innerWidth < 768) {
          t = false;
      }

      this.setState({
          ...this.state,
          showTaskListCols: t,
      })
  }

  _onOpenTasksHandler(res) {
      this.setState({
          ...this.state,
          nextOpenedTasksLink: res.data.links.next,
          prevOpenedTasksLink: res.data.links.prev,
          openedTasks: res.data.tasks.map(task => {
              return {id: task.id, name: task.name, done: task.done, description: task.description};
          }),
      });
  }

    _onClosedTasksHandler(res) {
        this.setState({
            ...this.state,
            nextClosedTasksLink: res.data.links.next,
            prevClosedTasksLink: res.data.links.prev,
            closedTasks: res.data.tasks.map(task => {
                return {id: task.id, name: task.name, done: task.done, description: task.description};
            }),
        });
    }

  _onErrorHandler(res) {
      alert(res);
  }

  _onGetOpenedTasks() {
      if(this.state.currentOpenedTasksUrl !== null
          && this.state.currentOpenedTasksUrl !== undefined
          && this.state.currentOpenedTasksUrl !== ""
          && this.state.isOpenedTasksTabOpen) {
          GetTasksByUrl(this.state.currentOpenedTasksUrl)
              .then(this._onOpenTasksHandler)
              .catch(this._onErrorHandler);
      } else {
          GetTasks(this.state.limit, 0, true)
              .then(this._onOpenTasksHandler)
              .catch(this._onErrorHandler);
      }
  }

  _onGetClosedTasks() {
      GetTasks(this.state.limit, 0, false)
          .then(this._onClosedTasksHandler)
          .catch(this._onErrorHandler);
  }

  componentWillMount () {
      this._onResize();

      // get the opened and closed tasks
      this._onGetOpenedTasks();
      this._onGetClosedTasks();
  }

  componentDidMount () {
      window.addEventListener("resize", this._onResize);
  }

  componentWillUnmount () {
      window.removeEventListener("resize", this._onResize);
  }

  render() {
      return (
          <Container className={"h-100"}>
              <Row>
                  <Col md="1" lg="2"/>
                  <Col xs="12" md="10" lg="8">
                      <InputGroup id="resume-container">
                          <Input placeholder="resume" innerRef={this.taskNameInput} onChange={this._onChangeText}/>
                          <InputGroupAddon addonType="append">
                              <Button disabled={!this.state.btnActive} onClick={this._onCreateTask}>Add</Button>
                          </InputGroupAddon>
                      </InputGroup>
                  </Col>
                  <Col md="1" lg="2"/>
                  <Col  md="1"  lg="2"/>
                  <Col xs="12" md="10" lg="8">
                      <InputGroup id="description-container">
                          <Input placeholder="description" innerRef={this.taskDescriptionInput} onChange={this._onChangeText}/>
                      </InputGroup>
                  </Col>
                  <Col md="1" lg="2"/>
              </Row>

              <Row className={"h-75"} id={"tasks"}>
                  {
                      this.state.showTaskListCols?
                          <Col xs="0" md="1" lg="2"/>:null
                  }
                  <Col xs="12" md="10" lg="8">
                      <TaskList
                          nextOpenedTasksLink={this.state.nextOpenedTasksLink}
                          prevOpenedTasksLink={this.state.prevOpenedTasksLink}
                          nextClosedTasksLink={this.state.nextClosedTasksLink}
                          prevClosedTasksLink={this.state.prevClosedTasksLink}
                          onTabChanged={this._onTabChanged}
                          onNext={this._onNextOrPrevt}
                          onPrev={this._onNextOrPrevt}
                      >
                          {
                              [...this.state.openedTasks, ...this.state.closedTasks].map((t, k) => {
                                  return <Task
                                      key={k}
                                      done={t.done}
                                      name={t.name}
                                      description={t.description}
                                      onDone={() => { this._onDone(t); }}
                                  />
                              })
                          }
                          </TaskList>
                  </Col>
                  {
                      this.state.showTaskListCols?
                          <Col xs="0" md="1" lg="2"/>:null
                  }
                  </Row>
          </Container>
      );
  }
}

export default App;
