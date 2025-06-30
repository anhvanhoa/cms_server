INSERT INTO mail_statuses (status, name, created_at) VALUES
('pending', 'Pending', NOW()),
('sent', 'Sent', NOW()),
('delivered', 'Delivered', NOW()),
('failed', 'Failed', NOW()),
('cancelled', 'Cancelled', NOW()),
('opened', 'Opened', NOW()),
('clicked', 'Clicked', NOW());

INSERT INTO type_mails (id, name, created_at) VALUES
('transactional', 'Transactional', NOW()),
('marketing', 'Marketing', NOW()),
('system', 'System', NOW()),
('notification', 'Notification', NOW());

INSERT INTO mail_providers (email, password, user_name, port, host, encryption, name, type_id, created_by, created_at, updated_at) VALUES
('toplaiphaiwin@gmail.com', 'ydhw abgh wwen mkqx', 'toplaiphaiwin@gmail.com', 587, 'smtp.gmail.com', 'tls', 'Cms Hệ Thống', 'system', null, NOW(), NOW());

INSERT INTO mail_templates (id, name, subject, body, keys, provider_email, status, created_by, created_at, updated_at) VALUES
('register_mail', 'Xác thực tải khoản hệ thống', 'Xác thực tài khoản', '<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
</head>
<body style="font-family: sans-serif; background-color: #f9f9f9; padding: 20px;">
  <div style="max-width: 600px; margin: auto; background-color: white; padding: 30px; border-radius: 10px;">
    
    <div style="text-align: center; margin: 20px 0;">
      <a href="{{link}}" style="
        display: inline-block;
        padding: 12px 24px;
        background-color: #28a745;
        color: white;
        text-decoration: none;
        border-radius: 6px;
        font-weight: bold;
      ">Xác thực tài khoản</a>
    </div>
    <p>Nếu bạn không tạo tài khoản, vui lòng bỏ qua email này.</p>
  </div>
</body>
</html>
', '{user, link}', 'toplaiphaiwin@gmail.com', 'active', null, NOW(), NOW()),
('forgot_mail', 'Láy lại mật khẩu', 'Lấy lại mật khẩu', '<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
</head>
<body style="font-family: sans-serif; background-color: #f9f9f9; padding: 20px;">
  <div style="max-width: 600px; margin: auto; background-color: white; padding: 30px; border-radius: 10px;">
    <div style="text-align: center; margin: 20px 0;">
    {{#link}}
      <a href="{{link}}" style="
        display: inline-block;
        padding: 12px 24px;
        background-color: #28a745;
        color: white;
        text-decoration: none;
        border-radius: 6px;
        font-weight: bold;
      ">Lấy lại mật khẩu</a>
    {{/link}}
    {{#code}}
      <p>
        Mã xác thực của bạn là: <strong>{{code}}</strong>
      </p>
    {{/code}}
    </div>
    <p>Nếu bạn không tạo tài khoản, vui lòng bỏ qua email này.</p>
  </div>
</body>
</html>
', '{user, link, code}', 'toplaiphaiwin@gmail.com', 'active', null, NOW(), NOW());