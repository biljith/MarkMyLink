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
                            src={this.props.image}
                        />
                        <Card.Header>{this.props.name}</Card.Header>
                        <Card.Meta>{this.props.link}</Card.Meta>
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