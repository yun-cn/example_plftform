-- user
create user if not exists         'noonde_api_server_w'@'%' identified by 'noonde_api_server_w';
grant  all  on noonde_api_server_development.* to 'noonde_api_server_w'@'%';

create user   if not exists         'noonde_api_server_r'@'%' identified by 'noonde_api_server_r';
grant  select on noonde_api_server_development.* to 'noonde_api_server_r'@'%';
