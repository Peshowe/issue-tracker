
import { Google } from '@styled-icons/boxicons-logos/Google';
import styled from 'styled-components'

const Icon = styled(Google)`
  color: #fff;
  transition: transform 200ms ease-in-out;
  &:hover {
    transform: scale(1.2);
  }
`


function Login(props) {

    // function loginRedirect() {
    //     props.history.push("/auth/google");
    // }

    return (
        <div>
            <h3>Login: </h3>
            <a
                href={window.location.origin + "/auth/google"}
            >
                <Icon
                    // onClick={loginRedirect}
                    size={36}
                    style={{
                        cursor: 'pointer',
                        textAlign: 'center'
                    }}
                />
            </a>
        </div>
    )
}

export default Login