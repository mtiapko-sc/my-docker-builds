<?php
$uri = $_SERVER['REQUEST_URI'] ?? '/';
if ($uri === '/healthz') {
    echo 'ok';
} else {
    echo 'Hello from PHP!';
}
