import React, { useState } from 'react'
import styled from 'styled-components'
import { LogOut } from '@styled-icons/boxicons-regular/LogOut'


const Icon = styled(LogOut)`
&:hover {
    transform: scale(1.2);
  }
`

function LogoutButton(props) {

    function logout() {
        fetch("/auth/logout")
            .then(
                (result) => {
                    window.location.reload();
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    window.location.reload();
                }
            )
    }

    return (
        <span style={{
            "float": "right",
            "paddingRight": "0.5em"
        }}>
            {props.user}
            <Icon
                size={48}
                style={{
                    borderRadius: '6px',
                    padding: '8px',
                    cursor: 'pointer',

                }}
                onClick={logout}
            />
        </span>
    )
}

export default LogoutButton