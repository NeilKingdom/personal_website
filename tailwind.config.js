/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ['index.html', './res/**/*.{html,js}'],
    theme: {
        extend: {
            scale: {
                '200': '2.0'
            },
            keyframes: {
                conveyorX: {
                    '0%': {
                        transform: 'translateX(-110%)',
                        opacity: '0'
                    },
                    '15%': {
                        opacity: '100'
                    },
                    '85%': {
                        opacity: '100'
                    },
                    '100%': {
                        transform: 'translateX(60vw)',
                        opacity: '0'
                    }
                }
            },
            animation: {
                'conveyor-x': 'conveyorX 20s linear infinite',
                'conveyor-x-rev': 'conveyorX reverse 20s linear infinite'
            }
        }
    },
    plugins: [
        require("tailwindcss-animation-delay")
    ],
}

