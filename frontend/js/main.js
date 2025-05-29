document.addEventListener('DOMContentLoaded', function() {
    fetch('/BookTalks-site/backend/check_auth.php', { credentials: 'same-origin' })
        .then(res => res.json())
        .then(data => {
            if (data.authenticated && data.username) {
                const welcome = document.createElement('div');
                welcome.className = 'welcome-user';
                welcome.textContent = `Привет, ${data.username}!`;
                const container = document.querySelector('.container');
                if (container) {
                    container.prepend(welcome);
                } else {
                    document.body.prepend(welcome);
                }
            }
        });
});
