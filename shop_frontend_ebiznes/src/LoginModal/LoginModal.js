import './LoginModal.css'
import axios from "axios";

const LoginModal = ({onClose})=>{

    const handleGithubLogin = () => {
        axios.get('https://shopworkingbackend.azurewebsites.net/oauth/login/github').then((data) => {
            window.open(data.data, '_self', 'noopener,noreferrer');
        })
    }
    const handleGoogleLogin = () => {
        axios.get('https://shopworkingbackend.azurewebsites.net/oauth/login/google').then((data) => {
            window.open(data.data, '_self', 'noopener,noreferrer');
        })
    }

    const handleFacebookLogin = () => {
        axios.get('https://shopworkingbackend.azurewebsites.net/oauth/login/facebook').then((data) => {
            window.open(data.data, '_self', 'noopener,noreferrer');
        })
    }

    const handleLinkedinLogin = () => {
        axios.get('https://shopworkingbackend.azurewebsites.net/oauth/login/linkedin').then((data) => {
            window.open(data.data, '_self', 'noopener,noreferrer');
        })
    }

    return (
        <div className="payment-modal">
            <button onClick={handleGithubLogin}>Github</button>
            <button onClick={handleGoogleLogin}>Google</button>
            <button onClick={handleFacebookLogin}>Facebook</button>
            <button onClick={handleLinkedinLogin}>Linkedin</button>
        </div>
    )
}

export default LoginModal;