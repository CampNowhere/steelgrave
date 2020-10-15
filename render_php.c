#include <php/20151012/sapi/embed/php_embed.h>

int render_php(char * php_file)
{
  zend_file_handle script;
  script.type = ZEND_HANDLE_FP;
  script.filename = php_file;
  script.opened_path = NULL;
  script.free_filename = 0;
  script.handle.fp = fopen(script.filename, "rb");
  if(!script.handle.fp)
  {
    printf("Unable to open %s\n", script.filename);
    return -1;
  }
  PHP_EMBED_START_BLOCK(1, &php_file);
    php_execute_script(&script TSRMLS_CC);
  PHP_EMBED_END_BLOCK();
  return 0;
}
