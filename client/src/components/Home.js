import { createMedia } from '@artsy/fresnel'
import PropTypes from 'prop-types'
import React, { Component, useState, useEffect } from 'react'
import {Card} from 'semantic-ui-react'
import Bookmark, {BookmarkModal} from './Bookmark.js'
import {
    Button,
    Container,
    Divider,
    Grid,
    Header,
    Icon,
    Image,
    List,
    Menu,
    Segment,
    Sidebar,
    Visibility,
    Label,
} from 'semantic-ui-react'

const { MediaContextProvider, Media } = createMedia({
    breakpoints: {
        mobile: 0,
        tablet: 768,
        computer: 1024,
    },
})

/* Heads up!
 * HomepageHeading uses inline styling, however it's not the best practice. Use CSS or styled
 * components for such things.
 */
const HomepageHeading = ({ mobile }) => (
    <Container text>
        <Header
            as='h1'
            content='Welcome to your bookmark manager'
            inverted
            style={{
                fontSize: mobile ? '2em' : '4em',
                fontWeight: 'normal',
                marginBottom: 0,
                marginTop: mobile ? '1.5em' : '3em',
            }}
        />
        <Header
            as='h2'
            content='Create, save , view and delete your bookmarks all in one place'
            inverted
            style={{
                fontSize: mobile ? '1.5em' : '1.7em',
                fontWeight: 'normal',
                marginTop: mobile ? '0.5em' : '1.5em',
            }}
        />
        <BookmarkModal/>
    </Container>
)

HomepageHeading.propTypes = {
    mobile: PropTypes.bool,
}

/* Heads up!
 * Neither Semantic UI nor Semantic UI React offer a responsive navbar, however, it can be implemented easily.
 * It can be more complicated, but you can create really flexible markup.
 */
class DesktopContainer extends Component {
    state = {}

    hideFixedMenu = () => this.setState({ fixed: false })
    showFixedMenu = () => this.setState({ fixed: true })

    logout = () => {
        localStorage.removeItem('token');
        window.location.reload(false);
    }

    render() {
        const { children } = this.props
        const { fixed } = this.state

        return (
            <Media greaterThan='mobile'>
                <Visibility
                    once={false}
                    onBottomPassed={this.showFixedMenu}
                    onBottomPassedReverse={this.hideFixedMenu}
                >
                    <Segment
                        inverted
                        textAlign='center'
                        style={{ minHeight: 700, padding: '1em 0em' }}
                        vertical
                    >
                        <Menu
                            fixed={fixed ? 'top' : null}
                            inverted={!fixed}
                            pointing={!fixed}
                            secondary={!fixed}
                            size='large'
                        >
                            <Container>
                                <Menu.Item as='a' active>
                                    Home
                                </Menu.Item>
                                <Menu.Item as='a'>Work</Menu.Item>
                                <Menu.Item as='a'>Company</Menu.Item>
                                <Menu.Item as='a'>Careers</Menu.Item>
                                <Menu.Item position='right' onClick={this.logout}>
                                    <Button as='a' inverted={!fixed}>
                                        Log Out
                                    </Button>
                                </Menu.Item>
                            </Container>
                        </Menu>
                        <HomepageHeading />
                    </Segment>
                </Visibility>

                {children}
            </Media>
        )
    }
}

DesktopContainer.propTypes = {
    children: PropTypes.node,
}

const ResponsiveContainer = ({ children }) => (
    /* Heads up!
     * For large applications it may not be best option to put all page into these containers at
     * they will be rendered twice for SSR.
     */
    <MediaContextProvider>
        <DesktopContainer>{children}</DesktopContainer>
    </MediaContextProvider>
)

ResponsiveContainer.propTypes = {
    children: PropTypes.node,
}

function HomepageLayout(props) { 
    const [bookmarks, setBookmarks] = useState([]);
    const [fetched, setFetched] = useState(false)
    console.log('Billy', props.token)
    const fetchBookmarks = async () => {
        var token = props.token
        await fetch('/bookmarks', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            token
          })
        }).then(response => response.json())
          .then(function(data) {
            console.log('Billy', data)
            setBookmarks(data)
            setFetched(true)
          },
          function(err) {

          });
    }
    if (!fetched)
        fetchBookmarks(props);
    
    return (
    <ResponsiveContainer>
        <Segment style={{ padding: '8em 0em' }} vertical>
            <Grid stackable verticalAlign='middle' style={{ paddingLeft: '7em' }}>
                <Grid.Row>
                    <Grid.Column>
                        <Card.Group>
                            {
                                bookmarks.map(function(bookmark) {
                                    return (
                                        <Bookmark name = { bookmark.Name} link={ bookmark.Link } image = {bookmark.Image} description = {bookmark.Description} viewcount = {bookmark.Viewcount}></Bookmark>
                                    );
                                })
                            }
                        </Card.Group>
                    </Grid.Column>
                </Grid.Row>
            </Grid>
        </Segment>


        <Segment inverted vertical style={{ padding: '5em 0em' }}>
            <div>
                <Header inverted as='h2' icon textAlign='center'>
                    <Icon  name='users' circular />
                    <Header.Content>Team Geauxinue</Header.Content>
                </Header>
            </div>

            <Grid divided inverted stackable style={{ paddingLeft: '17em' }} textAlign='center'>
                    <Grid.Row>
                        <Grid.Column>
                            <Card.Group>
                                <Card>
                                    <Card.Content>
                                        <Image
                                            floated='right'
                                            size='mini'
                                            src='https://react.semantic-ui.com/images/avatar/large/steve.jpg'
                                        />
                                        <Card.Header>Aakarsh Fadnis</Card.Header>
                                        <Card.Meta>Friends of Elliot</Card.Meta>
                                        <Card.Description>
                                            Steve wants to add you to the group <strong>best friends</strong>
                                        </Card.Description>
                                    </Card.Content>
                                </Card>
                                    <Card>
                                        <Card.Content>
                                            <Image
                                                floated='right'
                                                size='mini'
                                                src='https://react.semantic-ui.com/images/avatar/large/steve.jpg'
                                            />
                                            <Card.Header>Abhas Prasad</Card.Header>
                                            <Card.Meta>Friends of Elliot</Card.Meta>
                                            <Card.Description>
                                                Steve wants to add you to the group <strong>best friends</strong>
                                            </Card.Description>
                                        </Card.Content>
                                    </Card>
                                <Card>
                                    <Card.Content>
                                        <Image
                                            floated='right'
                                            size='mini'
                                            src='https://react.semantic-ui.com/images/avatar/large/steve.jpg'
                                        />
                                        <Card.Header>Biljith Thadichi</Card.Header>
                                        <Card.Meta>Friends of Elliot</Card.Meta>
                                        <Card.Description>
                                            Steve wants to add you to the group <strong>best friends</strong>
                                        </Card.Description>
                                    </Card.Content>

                                </Card>
                                <Card>
                                    <Card.Content>
                                        <Image
                                            floated='right'
                                            size='mini'
                                            src='https://react.semantic-ui.com/images/avatar/large/steve.jpg'
                                        />
                                        <Card.Header>Sid Keshkar</Card.Header>
                                        <Card.Meta>Friends of Elliot</Card.Meta>
                                        <Card.Description>
                                            Steve wants to add you to the group <strong>best friends</strong>
                                        </Card.Description>
                                    </Card.Content>

                                </Card>
                            </Card.Group>
                        </Grid.Column>
                    </Grid.Row>
                </Grid>

        </Segment>
    </ResponsiveContainer>
)
} 

export default HomepageLayout
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