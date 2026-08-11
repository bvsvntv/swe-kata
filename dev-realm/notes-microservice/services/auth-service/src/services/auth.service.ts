async function register() {
    console.log('register function @ auth service');

    await generateTokens();
}

async function login() {
    console.log('login function @ auth service');
}

async function logout() {
    console.log('logout function @ auth service');
}

async function generateTokens() {
    console.log('generateTokens function @ auth service');
}

async function refreshTokens() {
    console.log('refreshTokens function @ auth service');
}

export { register, login, logout, generateTokens, refreshTokens };
