import React, { Component } from 'react';
import {Card, Grid, Menu, Segment, Label, Button, Image, Icon} from 'semantic-ui-react'
import Bookmark from './Bookmark.js'

export default class Home extends Component{
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
                <Bookmark></Bookmark>

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