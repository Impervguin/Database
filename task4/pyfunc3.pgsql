-- Табличная функция

-- Выполняет транслитерацию имен и фамилий клиентов
-- DROP FUNCTION RuClient();
CREATE OR REPLACE FUNCTION RuClient ()
  RETURNS TABLE (cid int, first_name text, last_name text)
  LANGUAGE plpython3u
AS $$
eng_to_rus = {
    'a': 'а', 'b': 'б', 'v': 'в', 'g': 'г', 'd': 'д', 'e': 'е', 'yo': 'ё',
    'zh': 'ж', 'z': 'з', 'i': 'и', 'y': 'й', 'k': 'к', 'l': 'л', 'm': 'м',
    'n': 'н', 'o': 'о', 'p': 'п', 'r': 'р', 's': 'с', 't': 'т', 'u': 'у',
    'f': 'ф', 'kh': 'х', 'ts': 'ц', 'ch': 'ч', 'sh': 'ш', 'shch': 'щ',
    'y': 'ы', 'e': 'э', 'yu': 'ю', 'ya': 'я','j':'дж', 'h':'х'
}
def transliterate(txt):
  result = ""
  i = 0
  txt = txt.lower()
  while i < len(txt):
    if i < len(txt) - 1:
      double_letter = txt[i:i+2].lower()
      if double_letter in eng_to_rus:
        result += eng_to_rus[double_letter]
        i += 2
        continue
    letter = txt[i]
    result += eng_to_rus.get(letter, letter)
    i += 1
  return result
res = []
quer = "SELECT id, first_name, last_name FROM client"
cls = plpy.execute(quer)
for cl in cls:
	res.append((cl['id'], transliterate(cl['first_name']), transliterate(cl['last_name'])))
return res
$$;

SELECT * FROM RuClient();