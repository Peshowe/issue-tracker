import React, { useState } from 'react'

import { Menu as BurgerIcon } from '@styled-icons/remix-fill/Menu'
import { Menu, Overlay, StyledOffCanvas } from 'styled-off-canvas'

import List from './List'
import Close from './Close'
import GitHub from './GitHub'
import BackButton from './BackButton'
import LogoutButton from './LogoutButton'

import styled from 'styled-components'

const MenuIcon = styled(BurgerIcon)`
&:hover {
    transform: scale(1.2);
  }
`

function ToggleMenu(props) {
    const [isOpen, setIsOpen] = useState(false);

    let menu;
    let buttons;

    if (props.user == "") {
        menu = <Menu />
        buttons = <span></span>
    } else {
        menu = (<Menu>
            <>
                <Close onClose={() => setIsOpen(false)} />
                <List user={props.user} />
            </>
        </Menu>)
        buttons = (<span>
            <BackButton />
            <MenuIcon
                size={48}
                style={{
                    borderRadius: '6px',
                    padding: '8px',
                    cursor: 'pointer'

                }}
                onClick={() => { setIsOpen((isOpen) => !isOpen) }}
            />

            <GitHub />

            <LogoutButton user={props.user} />
        </span>)
    }

    return (
        <div>
            <StyledOffCanvas
                menuBackground='black'
                isOpen={isOpen}
                onClose={() => setIsOpen(false)}
            >
                {buttons}

                <span style={{
                    "padding": '8px',
                    "font-size": "1.7em",
                    "margin-left": "1em"
                }}>Parvus JIRA</span>

                {menu}


                <Overlay />
            </StyledOffCanvas>
        </div >
    )
}

export default ToggleMenu