<?php
// Подключение к базе данных
require_once 'db.php';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $authorFirstName = $_POST['authorFirstName'] ?? '';
    $authorLastName = $_POST['authorLastName'] ?? '';

    if (empty($authorFirstName) || empty($authorLastName)) {
        echo json_encode(['success' => false, 'message' => 'Имя и фамилия автора обязательны для заполнения.']);
        exit;
    }

    $stmt = $pdo->prepare("INSERT INTO author (first_name, last_name) VALUES (:first_name, :last_name)");
    $stmt->bindParam(':first_name', $authorFirstName);
    $stmt->bindParam(':last_name', $authorLastName);

    try {
        $stmt->execute();
        echo json_encode(['success' => true, 'message' => 'Автор успешно добавлен.']);
    } catch (PDOException $e) {
        echo json_encode(['success' => false, 'message' => 'Ошибка при добавлении автора: ' . $e->getMessage()]);
    }
} else {
    echo json_encode(['success' => false, 'message' => 'Неверный метод запроса.']);
}
?>