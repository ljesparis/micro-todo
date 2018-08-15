import React, { Component } from 'react';
import { Container, Row, Col, ListGroupItem, Button } from 'reactstrap';

import "./task.css";


class Task extends Component {

    _renderDoneButton() {
        if(this.props.done) {
            return null;
        } else {
            return (
                <Col xs="3" md="2" id="close-button-container">
                    <Button onClick={this.props.onDone} id="task-button-close">Done</Button>
                </Col>
            )
        }
    }

    render() {
        return (
            <ListGroupItem>
                <Container id="task-cotnainer">
                    <Row>
                        <Col xs={this.props.done?"12":"9"} md={this.props.done?"12":"10"} >
                            <Col>
                                <strong>{this.props.name}</strong>
                            </Col>
                            <Col>
                                {this.props.description}
                            </Col>
                        </Col>
                        {this._renderDoneButton()}
                    </Row>
                </Container>
            </ListGroupItem>
        )
    }
}


export default Task;
