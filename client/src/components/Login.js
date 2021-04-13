import React, {useState} from 'react'
import { Button, Form, Grid, Header, Message, Segment } from 'semantic-ui-react'
import { Redirect, useHistory, Link } from 'react-router-dom';

function Login ({setToken}) {
  const [email, setEmail] = useState();
  const [password, setPassword] = useState();
  const [error, setError] = useState(false);
  const history = useHistory();



  const handleSubmit = async e => {
    e.preventDefault();
    await fetch('/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
                            email,
                            password
                          })
    })
      .then(function(response) {
            if (!response.ok) {throw Error(response.statusText)}
            return response.json()
          }
      )
      .then(
        function(response) {
          const token =  response['token'];
          console.log(token);
          history.push("/")
          localStorage.setItem('token', token)
          setToken(token);
        },
        function(error) {
          console.log("error")
          setError(true);
        }
      )
      .catch(err => {setError(true);console.log("error")})
  };

  return (
    <Grid textAlign='center' style={{ height: '100vh' }} verticalAlign='middle'>
      <Grid.Column style={{ maxWidth: 450 }}>
        <Header as='h2' color='teal' textAlign='center'>
          Log-in to your account
        </Header>
        <Form size='large' onSubmit={handleSubmit} error={error}>
          <Segment stacked>
            <Form.Input fluid icon='user' iconPosition='left' placeholder='E-mail address' onChange={e => setEmail(e.target.value)} />
            <Form.Input
              fluid
              icon='lock'
              iconPosition='left'
              placeholder='Password'
              type='password'
              onChange={e => setPassword(e.target.value)}
            />

            <Button color='teal' fluid size='large'>
              Login
            </Button>
          </Segment>
          <Message
            error
            header='Login failed'>
              Login failed with an error.
          </Message>
        </Form>
        <Message>
          New to us? <Link to="/signup">Sign Up</Link>
        </Message>
      </Grid.Column>
    </Grid>
  );
}

export default Login