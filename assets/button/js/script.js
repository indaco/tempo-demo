// JS file for 'button'
function getButtonVariant(button) {
  return button.getAttribute('data-variant') || 'default'
}

function handleClick(event) {
  const variant = getButtonVariant(event)
  alert(`Variant: ${variant}`)
}
