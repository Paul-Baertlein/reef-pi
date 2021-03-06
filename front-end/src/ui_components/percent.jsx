import React from 'react'
import PropTypes from 'prop-types'

const Percent = props => {
  const { value, ...other } = props

  const handleChange = e => {
    // TODO: [ML] Allow decimal in regex
    if (/^([0-9]{0,2}$)|(100)$|^([0-9]{1,2}.[0-9]+$)/.test(e.target.value)) {
      let val = parseFloat(e.target.value)
      if (isNaN(e.target.value)) { val = '' }
      const event = {
        target: {
          name: e.target.name,
          value: val
        }
      }
      props.onChange(event)
    }
  }

  return (
    <input
      {...other}
      type='number'
      value={(isNaN(value) ? '' : value)}
      onChange={handleChange}
    />
  )
}

Percent.propTypes = {
  value: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
  onChange: PropTypes.func.isRequired
}

export default Percent
