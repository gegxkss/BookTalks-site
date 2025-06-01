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
                    document.getElementById('profile-nickname').textContent = profile.user.nickname || '';
                    document.getElementById('profile-email').textContent = profile.user.email || '';
                    document.getElementById('profile-name').textContent = profile.user.first_name || '';
                    document.getElementById('profile-lastname').textContent = profile.user.last_name || '';
                    document.getElementById('profile-sex').textContent = (profile.user.sex === 'male') ? 'мужской' : (profile.user.sex === 'female' ? 'женский' : '');
                    document.getElementById('profile-birthdate').textContent = profile.user.birth_date ? new Date(profile.user.birth_date).toLocaleDateString('ru-RU') : '';
                    if (profile.user.profile_image) {
    // Всегда абсолютный путь от корня сайта
    document.getElementById('profile-image').src = '/BookTalks-site/' + profile.user.profile_image.replace(/^\/+/, '');
} else {
    document.getElementById('profile-image').src = '/BookTalks-site/styles/images/icon.png';
}
                }
                // Добавляем вывод значения profile_image в консоль для отладки
                console.log('Profile image path:', profile.user.profile_image);
                // Можно дополнительно отобразить книги, рецензии, цитаты
                // ...
            }).catch(err => {
                console.error('Ошибка при разборе профиля:', err);
                // Можно вывести сообщение пользователю
            });
    });
