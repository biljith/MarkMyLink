import React, { Component } from 'react';
import {Card, Grid, Menu, Segment, Label, Button, Image, Icon, Modal, Form} from 'semantic-ui-react'

class AddBookmark extends Component {

    constructor(props) {
        super(props);
        this.state = {
            name: "",
            link: ""
        };
        this.handleFormSubmit = this.handleFormSubmit.bind(this);
    }
    

    handleFormSubmit(e) {
        e.preventDefault();
        var token = localStorage.getItem('token')
        var name = this.state.name
        var link = this.state.link
        fetch('/addBookmark', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            name,
            link,
            token
          })
        }).then(
            function(response) {
                window.location.reload(false)
            },
            function(error) {
                window.location.reload(false)
            }
        );
        // whatever you want to do when user submits a form
    };

    render() {
        const { fields } = this.state;

        return (
            <Form onSubmit={this.handleFormSubmit}>
                <Form.Field>
                  <label>Bookmark Name</label>
                  <input placeholder='Bookmark Name' onChange={e => this.setState({name: e.target.value})} />
                </Form.Field>
                <Form.Field>
                  <label>Link</label>
                  <input placeholder='Link' onChange={e => this.setState({link: e.target.value})} />
                </Form.Field>
                <Button type='submit'>Submit</Button>
            </Form>
        );
    }
}

class BookmarkModal extends Component {
    state = {
        modalOpen: false,
    };

    handleOpen = () => this.setState({ modalOpen: true });

    handleClose = () => this.setState({ modalOpen: false });

    render() {
        return (
            <div>
                <Button primary size='huge' onClick={this.handleOpen}>
                    Add Bookmark
                    <Icon name='right arrow' />
                </Button>
                <Modal
                    open={this.state.modalOpen}
                    onClose={this.handleClose}
                    closeIcon
                >
                    <Modal.Header>Create Bookmark</Modal.Header>
                    <Modal.Content>
                        <AddBookmark handleClose={this.handleClose} />
                    </Modal.Content>
                </Modal>
            </div>
        )
    }
}

export default class Bookmark extends Component{
	constructor(props) {
		super(props);
	}
    render() {
        return (
                <Card>
                    <Label color="teal" floating>
                        {this.props.viewcount}
                    </Label>
                    <Card.Content>
                        <Image
                            floated='right'
                            size='tiny'
                            src={this.props.image}
                        />
                        <Card.Header>{this.props.name}</Card.Header>
                        <Card.Meta>{this.props.link}</Card.Meta>
                        <Label as='a' style={{backgroundColor:'#2185d0', color:'white'}} ribbon>
                            Overview
                        </Label>
                        <Card.Description>
                            {this.props.description}
                        </Card.Description>
                    </Card.Content>
                    <Card.Content extra centered>
                        {/*<div className='ui three buttons' textAlign="center">*/}
                        <Button.Group widths={'3'}>
                            <Button floated='left' animated={"fade"} as='a' href={'https://'+this.props.link}>
                                <Button.Content hidden>View</Button.Content>
                                <Button.Content visible>
                                    <Icon name='eye' />
                                </Button.Content>
                            </Button>
                            <Button animated={"fade"}>
                                <Button.Content hidden>Edit</Button.Content>
                                <Button.Content visible>
                                    <Icon name='edit' />
                                </Button.Content>
                            </Button>
                            <Button floated='left' animated={"fade"} onClick={this.handleDelete}>
                                <Button.Content hidden>Delete</Button.Content>
                                <Button.Content visible>
                                    <Icon name='delete' />
                                </Button.Content>
                            </Button>
                        </Button.Group>
                        {/*</div>*/}
                    </Card.Content>
                </Card>
        )
    }
}

export {
    BookmarkModal
}