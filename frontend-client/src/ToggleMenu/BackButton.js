import React, { useState } from 'react'
// import { useHistory } from "react-router-dom";
import styled from 'styled-components'
import { ArrowBack } from '@styled-icons/boxicons-regular/ArrowBack'


const Icon = styled(ArrowBack)`
&:hover {
    transform: scale(1.2);
  }
`

function BackButton(props) {

    // const history = useHistory();

    function goBack() {
        // history.goBack();
        if (window.location.pathname != "/") {
            window.history.back();
        }
    }

    return (<Icon
        size={48}
        style={{
            borderRadius: '6px',
            padding: '8px',
            cursor: 'pointer',

        }}
        onClick={goBack}
    />
    )
}

export default BackButton