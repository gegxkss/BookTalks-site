document.addEventListener('DOMContentLoaded', function() {
    const registerForm = document.getElementById('registerForm');
    const notification = document.getElementById('notification');

    function showNotification(message, type = 'error') {
        notification.textContent = message;
        notification.className = `notification ${type}`;
        notification.classList.add('active');
        
        setTimeout(() => {
            notification.classList.remove('active');
        }, 3000);
    }

    registerForm.addEventListener('submit', async function(event) {
        event.preventDefault();

        const formData = new FormData(registerForm);
        const submitButton = registerForm.querySelector('button[type="submit"]');
        const originalButtonText = submitButton.textContent;

        try {
            submitButton.disabled = true;
            submitButton.innerHTML = '<span class="loading-spinner"></span>';

            const response = await fetch('/api/register', {
                method: 'POST',
                body: formData
            });

            const result = await response.json();

            if (response.ok) {
                showNotification('Регистрация успешна! Перенаправление...', 'success');
                setTimeout(() => {
                    window.location.href = '/';
                }, 1000);
            } else {
                showNotification(result.message || 'Ошибка регистрации');
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