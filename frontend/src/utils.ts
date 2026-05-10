export const fmtDate = (v:string)=> new Intl.DateTimeFormat('ru-RU',{day:'2-digit',month:'short',hour:'2-digit',minute:'2-digit'}).format(new Date(v))
export const timeAgo = (v:string)=> { const diff=Math.floor((Date.now()-+new Date(v))/3600000); return diff<1?'только что':diff<24?`${diff} ч назад`:`${Math.floor(diff/24)} дн назад` }
export const uid = (p:string)=> `${p}-${Math.random().toString(36).slice(2,8)}`
export const mdToHtml = (s:string)=> s
  .replace(/^## (.*)$/gm,'<h2>$1</h2>')
  .replace(/^1\. (.*)$/gm,'<li>$1</li>')
  .replace(/```([\s\S]*?)```/g,'<pre class="code"><code>$1</code></pre>')
  .replace(/\n\n/g,'</p><p>')
  .replace(/^/,'<p>')
  .replace(/$/,'</p>')
