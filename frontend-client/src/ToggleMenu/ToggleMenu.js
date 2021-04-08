import React, { useState } from 'react'
import { Menu as BurgerIcon } from '@styled-icons/remix-fill/Menu'

import { Menu, Overlay, StyledOffCanvas } from 'styled-off-canvas'

import List from './List'
import Close from './Close'
import GitHub from './GitHub'

function ToggleMenu() {
    const [isOpen, setIsOpen] = useState(false)

    return (
        <div>
            <StyledOffCanvas
                menuBackground='black'
                isOpen={isOpen}
                onClose={() => setIsOpen(false)}
            >
                <BurgerIcon
                    size={48}
                    style={{
                        borderRadius: '6px',
                        padding: '8px',
                        cursor: 'pointer'

                    }}
                    onClick={() => { setIsOpen((isOpen) => !isOpen) }}
                />

                <GitHub />

                <Menu>
                    <>
                        <Close onClose={() => setIsOpen(false)} />
                        <List />
                    </>
                </Menu>

                <Overlay />
            </StyledOffCanvas>
        </div>
    )
}

export default ToggleMenu