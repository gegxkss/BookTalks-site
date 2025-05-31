// Проверка авторизации и загрузка профиля пользователя
fetch('/BookTalks-site/backend/check_auth.php', { credentials: 'same-origin' })
    .then(res => res.json())
    .then(data => {
        if (!data.authenticated) {
            window.location.href = '/BookTalks-site/frontend/signIn.html';
            return;
        }
        // Если авторизован, загружаем данные профиля
        fetch('/BookTalks-site/backend/user_profile_data.php', { credentials: 'same-origin' })
            .then(res => res.json())
            .then(profile => {
                if (profile.user) {
                    document.getElementById('profile-nickname').textContent = profile.user.nickname;
                    document.getElementById('profile-email').textContent = profile.user.email;
                    document.getElementById('profile-name').textContent = profile.user.first_name || '';
                    document.getElementById('profile-lastname').textContent = profile.user.last_name || '';
                    document.getElementById('profile-sex').textContent = profile.user.sex || '';
                    document.getElementById('profile-birthdate').textContent = profile.user.birth_date || '';
                    if (profile.user.profile_image) {
                        document.getElementById('profile-image').src = '/BookTalks-site/frontend/' + profile.user.profile_image;
                    }
                }
                // Можно дополнительно отобразить книги, рецензии, цитаты
                // ...
            });
    });
