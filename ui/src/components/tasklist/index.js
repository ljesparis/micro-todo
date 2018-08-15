import React, { Component } from 'react';
import {
    ListGroup,
    Card,
    CardBody,
    Nav,
    NavItem,
    NavLink,
    TabPane,
    TabContent,
    CardFooter,
    Button
} from 'reactstrap';

import classnames from 'classnames';

import "./tasklist.css"

class TaskList extends Component {

    constructor(props) {
        super(props);

        this.nextRef = React.createRef();
        this.prevRef = React.createRef();

        this.toggle = this.toggle.bind(this);
        this.state = {
            activeTab: '1',
        };
    }

    toggle(tab) {
        if (this.state.activeTab !== tab) {
            this.setState({
                activeTab: tab
            });

            if(this.props.onTabChanged !== null
                && this.props.onTabChanged !== undefined
                && typeof this.props.onTabChanged === "function") {
                this.props.onTabChanged(tab);
            }
        }
    }

    render() {
        const {
            nextOpenedTasksLink,
            prevOpenedTasksLink,
            nextClosedTasksLink,
            prevClosedTasksLink
        } = this.props;

        let nextUrl = "";
        let prevUrl = "";

        if(this.state.activeTab === '1') {
            nextUrl = nextOpenedTasksLink;
            prevUrl = prevOpenedTasksLink;
        } else {
            nextUrl = nextClosedTasksLink;
            prevUrl = prevClosedTasksLink;
        }

        return (
            <Card className={"h-100"}>
                <CardBody>
                    <Nav tabs>
                        <NavItem>
                            <NavLink
                                className={classnames({ active: this.state.activeTab === '1' })}
                                onClick={() => { this.toggle('1'); }}
                            >
                                Open
                            </NavLink>
                        </NavItem>
                        <NavItem>
                            <NavLink
                                className={classnames({ active: this.state.activeTab === '2' })}
                                onClick={() => { this.toggle('2'); }}
                            >
                                Closed
                            </NavLink>
                        </NavItem>
                    </Nav>
                    <TabContent activeTab={this.state.activeTab}>
                        <TabPane tabId="1">
                            <ListGroup>
                                {this.props.children.map(t => {
                                    return !t.props.done? t:null;
                                })}
                            </ListGroup>
                        </TabPane>
                        <TabPane tabId="2">
                            <ListGroup>
                                {this.props.children.map(t => {
                                    return t.props.done? t:null;
                                })}
                            </ListGroup>
                        </TabPane>
                    </TabContent>
                </CardBody>
                <CardFooter>
                    <Button id="prev" innerRef={this.prevRef}
                            disabled={prevUrl === null || prevUrl === undefined || prevUrl.length === 0 }
                            onClick={() => {
                                this.props.onPrev(prevUrl);
                            }}
                    >
                        Prev
                    </Button>
                    <Button id="next" innerRef={this.nextRef}
                            disabled={nextUrl === null || nextUrl === undefined || nextUrl.length === 0 }
                            onClick={() => {
                                this.props.onNext(nextUrl);
                            }}
                    >
                        Next
                    </Button>
                </CardFooter>
            </Card>
        )
    }
}


export default TaskList;
