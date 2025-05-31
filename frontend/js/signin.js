document.addEventListener('DOMContentLoaded', function() {
    const loginForm = document.getElementById('loginForm');
    if (!loginForm) return;
    
    loginForm.addEventListener('submit', function(event) {
        event.preventDefault();
        const email = document.getElementById('email').value.trim();
        const password = document.getElementById('password').value;
        document.getElementById('emailError').textContent = '';
        document.getElementById('passwordError').textContent = '';
        document.getElementById('loginError').textContent = '';

        let valid = true;
        if (!email) {
            document.getElementById('emailError').textContent = 'Введите почту или ник';
            valid = false;
        }
        if (!password) {
            document.getElementById('passwordError').textContent = 'Введите пароль';
            valid = false;
        }
        if (!valid) return;

        fetch('/BookTalks-site/backend/login.php', {
            method: 'POST',
            credentials: 'same-origin',
            headers: { 'Content-Type': 'application/x-www-form-urlencoded' },
            body: new URLSearchParams({ email, password })
        })
        .then(res => res.json())
        .then(data => {
            if (data.success) {
                window.location.href = '/BookTalks-site/frontend/profile.html';
            } else {
                document.getElementById('loginError').textContent = data.error || 'Ошибка входа';
                // Если ошибка авторизации, предложить регистрацию
                if (data.error && data.error.includes('Неверная')) {
                    const regLink = document.createElement('a');
                    regLink.href = '/BookTalks-site/frontend/registration.html';
                    regLink.textContent = 'Зарегистрироваться';
                    regLink.style.marginLeft = '10px';
                    document.getElementById('loginError').appendChild(regLink);
                }
            }
        })
        .catch(() => {
            document.getElementById('loginError').textContent = 'Ошибка соединения с сервером';
        });
    });
});