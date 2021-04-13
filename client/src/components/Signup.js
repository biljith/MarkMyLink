import React, {useState} from 'react'
import { Button, Form, Grid, Header, Message, Segment } from 'semantic-ui-react'
import { Redirect, Link } from 'react-router-dom';


function Signup () {
  const [email, setEmail] = useState();
  const [password, setPassword] = useState();
  const [name, setUsername] = useState();
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);


  
  const handleSubmit = async e => {
    e.preventDefault();
    await fetch('/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        name,
        email,
        password
      })
    }).then(
          function(response) {
            if (response.status != 200) {
              setError(true);
              setSuccess(false);
            }
            else {
              setError(false);
              setSuccess(true);
            }
          },
          function(error) {
            setSuccess(false);
            setError(true);
          }
      );
  }

  return (
    <Grid textAlign='center' style={{ height: '100vh' }} verticalAlign='middle'>
      <Grid.Column style={{ maxWidth: 450 }}>
        <Header as='h2' color='teal' textAlign='center'>
          Create your account
        </Header>
        <Form size='large' onSubmit={handleSubmit} success={success} error={error}>
          <Segment stacked>
            <Form.Input fluid icon='id badge' iconPosition='left' placeholder='Username' onChange={e => setUsername(e.target.value)} />
            <Form.Input fluid icon='user' iconPosition='left' placeholder='E-mail address' onChange={e => setEmail(e.target.value)} />
            <Form.Input
              fluid
              icon='lock'
              iconPosition='left'
              placeholder='Password'
              type='password'
              onChange={e => setPassword(e.target.value)}
            />
            <Message
              success
              header='Signup Complete'>
              Successfully signed up! <Link to="/login">Log In</Link>.
            </Message>
            <Message
              error
              header='Signup failed'>
              User with the same email already exists. <Link to="/login">Log In</Link>&nbsp;instead
            </Message>
            <Button color='teal' fluid size='large'> 
              Signup
            </Button>
          </Segment>
        </Form>
        <Message>
          Already have an account with us? <Link to="/login">Log In</Link>
        </Message>
      </Grid.Column>
    </Grid>
  )
}

export default Signup