import React from 'react'

import { Close as CloseIcon } from '@styled-icons/remix-fill/Close'

export default ({ onClose }) => (
    <div
        style={{
            padding: '10px 10px 20px 10px',
            textAlign: 'right'
        }}
    >
        <CloseIcon
            onClick={onClose}
            size={36}
            style={{
                cursor: 'pointer',
                textAlign: 'center'
            }}
        />
    </div>
)