import React, { Component } from 'react';
import {Card, Grid, Menu, Segment, Label, Button, Image, Icon} from 'semantic-ui-react'

const items = [
    {
        header: 'Project Report - April',
        description:
            'Leverage agile frameworks to provide a robust synopsis for high level overviews.',
        meta: 'ROI: 30%',
    },
    {
        header: 'Project Report - May',
        description:
            'Bring to the table win-win survival strategies to ensure proactive domination.',
        meta: 'ROI: 34%',
    },
    {
        header: 'Project Report - April',
        description:
            'Leverage agile frameworks to provide a robust synopsis for high level overviews.',
        meta: 'ROI: 30%',
    },
    {
        header: 'Project Report - May',
        description:
            'Bring to the table win-win survival strategies to ensure proactive domination.',
        meta: 'ROI: 34%',
    },
    {
        header: 'Project Report - April',
        description:
            'Leverage agile frameworks to provide a robust synopsis for high level overviews.',
        meta: 'ROI: 30%',
    },
    {
        header: 'Project Report - May',
        description:
            'Bring to the table win-win survival strategies to ensure proactive domination.',
        meta: 'ROI: 34%',
    },
    {
        header: 'Project Report - April',
        description:
            'Leverage agile frameworks to provide a robust synopsis for high level overviews.',
        meta: 'ROI: 30%',
    },
    {
        header: 'Project Report - May',
        description:
            'Bring to the table win-win survival strategies to ensure proactive domination.',
        meta: 'ROI: 34%',
    },
]

export default class Bookmark extends Component{
    state = {}

    handleItemClick = (e, { name }) => this.setState({ activeItem: name })
    render() {
        const { activeItem } = this.state

        return (
            <div>
            <Menu pointing inverted>
                <Menu.Item
                    name='Home'
                    active={activeItem === 'Home'}
                    onClick={this.handleItemClick}
                >
                    Home
                    <Label color='teal'>10</Label>
                </Menu.Item>

                <Menu.Item
                    name='Cat1'
                    active={activeItem === 'Cat1'}
                    onClick={this.handleItemClick}
                >

                    Category 1
                    <Label color='teal'>1</Label>
                </Menu.Item>

                <Menu.Item
                    name='Cat2'
                    active={activeItem === 'Cat2'}
                    onClick={this.handleItemClick}
                >

                    Category 2
                    <Label color='teal'>1</Label>
                </Menu.Item>
                <Menu.Menu position='right'>
                    <Menu.Item
                        name='logout'
                        active={activeItem === 'logout'}
                        onClick={this.handleItemClick}
                    />
                </Menu.Menu>
            </Menu>

            <Segment>
                {/*<Card.Group centered items={items}/>*/}
                <Card.Group>
                    <Card>
                        <Card.Content>
                            <Image
                                floated='right'
                                size='mini'
                                src='/images/avatar/large/steve.jpg'
                            />
                            <Card.Header>Steve Sanders</Card.Header>
                            <Card.Meta>Friends of Elliot</Card.Meta>
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
                    <Card>
                        <Card.Content>
                            <Image
                                floated='right'
                                size='mini'
                                src='/images/avatar/large/molly.png'
                            />
                            <Card.Header>Molly Thomas</Card.Header>
                            <Card.Meta>New User</Card.Meta>
                            <Card.Description>
                                Molly wants to add you to the group <strong>musicians</strong>
                            </Card.Description>
                        </Card.Content>
                        <Card.Content extra>
                            <div className='ui three buttons' textAlign="center">
                                <Button.Group centered inverted>
                                    <Button>View</Button>
                                    <Button>Edit</Button>
                                    <Button>Delete</Button>
                                </Button.Group>
                            </div>
                        </Card.Content>
                    </Card>
                    <Card>
                        <Card.Content>
                            <Image
                                floated='right'
                                size='mini'
                                src='/images/avatar/large/jenny.jpg'
                            />
                            <Card.Header>Jenny Lawrence</Card.Header>
                            <Card.Meta>New User</Card.Meta>
                            <Card.Description>
                                Jenny requested permission to view your contact details
                            </Card.Description>
                        </Card.Content>
                        <Card.Content extra>
                            <div className='ui three buttons' textAlign="center">
                                <Button.Group centered inverted>
                                    <Button>View</Button>
                                    <Button>Edit</Button>
                                    <Button>Delete</Button>
                                </Button.Group>
                            </div>
                        </Card.Content>
                    </Card>
                </Card.Group>

            </Segment>
            </div>
        )
    }
}
// import React from 'react'
// import { Card } from 'semantic-ui-react'
//
// const Bookmark = () => (
//     <Card.Group>
//       <Card fluid color='red' header='Option 1' />
//       <Card fluid color='orange' header='Option 2' />
//       <Card fluid color='yellow' header='Option 3' />
//     </Card.Group>
// )
//
// export default Bookmark