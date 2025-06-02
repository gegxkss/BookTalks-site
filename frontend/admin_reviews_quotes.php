<?php
// admin_reviews_quotes.php — простая админка для рецензий и цитат
require_once '../backend/db.php';

// Удаление рецензии
if (isset($_GET['delete_review'])) {
    $id = intval($_GET['delete_review']);
    $pdo->prepare('DELETE FROM review WHERE id = ?')->execute([$id]);
    header('Location: admin_reviews_quotes.php');
    exit;
}
// Удаление цитаты
if (isset($_GET['delete_quote'])) {
    $id = intval($_GET['delete_quote']);
    $pdo->prepare('DELETE FROM quote WHERE id = ?')->execute([$id]);
    header('Location: admin_reviews_quotes.php');
    exit;
}

$reviews = $pdo->query('SELECT r.id, u.nickname, b.name AS book, r.text, r.created_at FROM review r JOIN users u ON r.user_id = u.id JOIN book b ON r.book_id = b.id ORDER BY r.created_at DESC')->fetchAll();
$quotes = $pdo->query('SELECT q.id, u.nickname, b.name AS book, q.text FROM quote q JOIN users u ON q.user_id = u.id JOIN book b ON q.book_id = b.id ORDER BY q.id DESC')->fetchAll();
?>
<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <title>Админка: Рецензии и Цитаты</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        table { border-collapse: collapse; width: 100%; margin-bottom: 40px; }
        th, td { border: 1px solid #ccc; padding: 8px; }
        th { background: #f1efe3; }
        .del { color: #fff; background: #c00; border: none; padding: 4px 10px; border-radius: 4px; cursor: pointer; }
    </style>
</head>
<body>
    <h1>Рецензии</h1>
    <table>
        <tr><th>ID</th><th>Пользователь</th><th>Книга</th><th>Текст</th><th>Дата</th><th>Удалить</th></tr>
        <?php foreach ($reviews as $r): ?>
        <tr>
            <td><?= $r['id'] ?></td>
            <td><?= htmlspecialchars($r['nickname']) ?></td>
            <td><?= htmlspecialchars($r['book']) ?></td>
            <td><?= nl2br(htmlspecialchars($r['text'])) ?></td>
            <td><?= $r['created_at'] ?></td>
            <td><a href="?delete_review=<?= $r['id'] ?>" onclick="return confirm('Удалить рецензию?')"><button class="del">Удалить</button></a></td>
        </tr>
        <?php endforeach; ?>
    </table>

    <h1>Цитаты</h1>
    <table>
        <tr><th>ID</th><th>Пользователь</th><th>Книга</th><th>Текст</th><th>Удалить</th></tr>
        <?php foreach ($quotes as $q): ?>
        <tr>
            <td><?= $q['id'] ?></td>
            <td><?= htmlspecialchars($q['nickname']) ?></td>
            <td><?= htmlspecialchars($q['book']) ?></td>
            <td><?= nl2br(htmlspecialchars($q['text'])) ?></td>
            <td><a href="?delete_quote=<?= $q['id'] ?>" onclick="return confirm('Удалить цитату?')"><button class="del">Удалить</button></a></td>
        </tr>
        <?php endforeach; ?>
    </table>
</body>
</html>
