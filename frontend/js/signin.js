document.addEventListener('DOMContentLoaded', function() {
    const loginForm = document.getElementById('loginForm');
    const notification = document.getElementById('notification');

    function showNotification(message, type = 'error') {
        notification.textContent = message;
        notification.className = `notification ${type}`;
        notification.classList.add('active');
        
        setTimeout(() => {
            notification.classList.remove('active');
        }, 3000);
    }

    loginForm.addEventListener('submit', async function(event) {
        event.preventDefault();

        const formData = new FormData(loginForm);
        const submitButton = loginForm.querySelector('button[type="submit"]');
        const originalButtonText = submitButton.textContent;

        try {
            submitButton.disabled = true;
            submitButton.innerHTML = '<span class="loading-spinner"></span>';

            const response = await fetch('/api/login', {
                method: 'POST',
                body: formData
            });

            const result = await response.json();

            if (response.ok) {
                showNotification('Успешный вход! Перенаправление...', 'success');
                setTimeout(() => {
                    window.location.href = '/';
                }, 1000);
            } else {
                showNotification(result.message || 'Ошибка авторизации');
            }
        } catch (error) {
            console.error('Ошибка:', error);
            showNotification('Ошибка соединения с сервером');
        } finally {
            submitButton.disabled = false;
            submitButton.textContent = originalButtonText;
        }
    });
});