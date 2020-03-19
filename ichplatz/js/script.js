document.addEventListener('DOMContentLoaded', () => {
  window.fetch('coronaZahlen.json')
    .then(r => r.json())
    .then(render)
})

function renderTable (rows) {
  return rows.map(row => `
    <tr>
        <td><a href="${row.URL}">${row.Name}</a></td>
        <td class="${row.Count < row.Max ? 'blass' : ''}">${isAvailable(row.Count)}</td>
        <td class="${row.RKI < row.Max ? 'blass' : ''}">${row.RKI}</td>
        <td class="${row.Mopo < row.Max ? 'blass' : ''}">${row.Mopo}</td>
        <td>${timestamp(row.Date)}</td>
    </tr>
  `).join('\n')
}

function render (data) {
  if (window.data === undefined) {
    window.data = data
  }
  window.data.Regions = Object.entries(data.Regions).map(([Name, region], i) => ({ Index: i, Name, ...region }))
  window.data.Mopo.Count = window.data.Regions.map(r => r.Mopo).reduce((a, c) => a + c)
  document.querySelector('body').innerHTML = `
    <div class="container">
        <h1>Fallzahlen Corona Deutschlandweit</h1>

        <p>
            Die Fallzahlen sind aus den Homepages (z.T. Pressemitteilungen) der Bundesländer zusammengetragen.
            Die Gesamtzahl ist die Summe der für jedes Bundesland jeweils höheren Fallzahl (<abbr title="Robert-Koch-Institut">RKI</abbr>
            bzw. Bundesland) in der Hoffnung, damit die aktuellsten Zahl anzuzeigen. (Stand:&nbsp;${timestamp(data.Date)})
        </p>
        <p>
            Diese Website ist ein <a href="https://twitter.com/Hoffmann">privates</a> Projekt. Der Quellcode steht unter der GPLv3 und ist auf
            <a href="https://github.com/HoffmannP/coronaZahlen">Github</a> zu finden, ebenso wie die
            <a href="https://github.com/HoffmannP/coronaZahlen/releases">Binary</a>. Die Rohdaten sind separat im json-Format
            <a href="${data.Name}.json">downloadbar</a>. Kontakt via <a href="https://twitter.com/Hoffmann">Twitter <span class="tt">@Hoffmann</span></a>.
        </p>
        
        <table class="u-full-width">
            <thead>
                <tr>
                    <th class="sorted">Bundesland</th>
                    <th>Fälle</th>
                    <th>RKI</th>
                    <th>MoPo</th>
                    <th>Stand</th>
                </tr>
            </thead>
            <tbody>
                ${renderTable(data.Regions)}
            </tbody>
            <tfoot>
                <tr>
                    <th rowspan="3">Deutschland</th>
                    <th>${data.Sum}</th>
                    <td colspan="3"></td>
                </tr>
                <tr>
                    <td></td>
                    <td>${data.RKI.Count}</td>
                    <td></td>
                    <td><a target="_blank" href="${data.RKI.URL}">${timestamp(data.RKI.Date)}</a></td>
                </tr>
                <tr>
                    <td colspan="2"></td>
                    <td>${data.Mopo.Count}</td>
                    <td><a target="_blank" href="${data.Mopo.URL}">${timestamp(data.Mopo.Date)}</a></td>
                </tr>
            </tfoot>
        </table>
    </div>`
  window.th = document.querySelectorAll('thead > tr > th')
  window.th.forEach(
    (th, i) => th.addEventListener('click', sort.bind(0, i))
  )
  window.th[0].dataset.sort = 0
}

const sorter = [
  (a, b) => a.Index - b.Index,
  (b, a) => a.Index - b.Index,
  (a, b) => a.Count - b.Count,
  (b, a) => a.Count - b.Count,
  (a, b) => a.RKI - b.RKI,
  (b, a) => a.RKI - b.RKI,
  (a, b) => a.Mopo - b.Mopo,
  (b, a) => a.Mopo - b.Mopo,
  (a, b) => a.Date - b.Date,
  (b, a) => a.Date - b.Date
]

function sort (col) {
  window.th.forEach(th => th.classList.remove('sorted'))
  const th = window.th[col]
  th.classList.add('sorted')
  th.dataset.sort = +!+th.dataset.sort
  document.querySelector('tbody').innerHTML = renderTable(
    window.data.Regions.sort(sorter[2 * col + +th.dataset.sort])
  )
}

function timestamp (unix) {
  if (unix === 0) {
    return 'n/a'
  }
  const ts = new Date(unix * 1000)
  const withTime = ts.getHours() > 0 || ts.getMinutes() > 0
  const date = { day: 'numeric', month: '2-digit', year: 'numeric' }
  const time = { hour: 'numeric', minute: '2-digit' }
  return ts.toLocaleString('de-DE', { ...date, ...(withTime ? time : {}) }) + (withTime ? ' Uhr' : '')
}
const isAvailable = x => x === -1 ? 'n/a' : x
