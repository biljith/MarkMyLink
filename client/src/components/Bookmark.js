import React, { Component } from 'react';
import {Card, Grid, Menu, Segment, Label, Button, Image, Icon} from 'semantic-ui-react'

export default class Bookmark extends Component{
	constructor(props) {
		super(props);
	}
    render() {
        return (
                <Card>
                    <Card.Content>
                        <Image
                            floated='right'
                            size='mini'
                            src='/images/avatar/large/steve.jpg'
                        />
                        <Card.Header>{this.props.name}</Card.Header>
                        <Card.Meta>{this.props.link}</Card.Meta>
                        <Card.Description>
                            Steve wants to add you to the group <strong>best friends</strong>
                        </Card.Description>
                    </Card.Content>
                    <Card.Content extra centered>
                        <div className='ui three buttons' textAlign="center">
                            <Button.Group centered>
                                <Button >View</Button>
                                <Button>Edit</Button>
                                <Button>Delete</Button>
                            </Button.Group>
                        </div>
                    </Card.Content>
                </Card>
        )
    }
}