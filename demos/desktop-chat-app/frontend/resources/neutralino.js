/* eslint-disable */
let Neutralino
;(() => {
    const e = {
        885: (e, t, i) => {
            Object.defineProperty(t, '__esModule', { value: !0 }),
                (t.open =
                    t.getConfig =
                    t.keepAlive =
                    t.killProcess =
                    t.exit =
                        void 0)
            const n = i(69)
            ;(t.exit = function (e) {
                return n.request({
                    url: 'app.exit',
                    type: n.RequestType.POST,
                    data: { code: e },
                    isNativeMethod: !0,
                })
            }),
                (t.killProcess = function () {
                    return n.request({
                        url: 'app.killProcess',
                        type: n.RequestType.GET,
                        isNativeMethod: !0,
                    })
                }),
                (t.keepAlive = function () {
                    return n.request({
                        url: 'app.keepAlive',
                        type: n.RequestType.GET,
                        isNativeMethod: !0,
                    })
                }),
                (t.getConfig = function () {
                    return n.request({
                        url: 'app.getConfig',
                        type: n.RequestType.GET,
                        isNativeMethod: !0,
                    })
                }),
                (t.open = function (e) {
                    return n.request({
                        url: 'app.open',
                        type: n.RequestType.POST,
                        data: e,
                        isNativeMethod: !0,
                    })
                })
        },
        308: (e, t, i) => {
            Object.defineProperty(t, '__esModule', { value: !0 }),
                (t.getRamUsage = void 0)
            const n = i(69)
            t.getRamUsage = function () {
                return n.request({
                    url: 'computer.getRamUsage',
                    type: n.RequestType.GET,
                    isNativeMethod: !0,
                })
            }
        },
        199: (e, t, i) => {
            Object.defineProperty(t, '__esModule', { value: !0 }),
                (t.log = t.LoggerType = void 0)
            const n = i(69)
            let r
            ;((r = t.LoggerType || (t.LoggerType = {})).WARN = 'WARN'),
                (r.ERROR = 'ERROR'),
                (r.INFO = 'INFO'),
                (t.log = function (e) {
                    return n.request({
                        url: 'debug.log',
                        type: n.RequestType.POST,
                        data: e,
                        isNativeMethod: !0,
                    })
                })
        },
        284: (e, t) => {
            Object.defineProperty(t, '__esModule', { value: !0 }),
                (t.dispatch = t.off = t.on = void 0),
                (t.on = function (e, t) {
                    return new Promise((i, n) => {
                        window.addEventListener(e, t), i()
                    })
                }),
                (t.off = function (e, t) {
                    return new Promise((i, n) => {
                        window.removeEventListener(e, t), i()
                    })
                }),
                (t.dispatch = function (e, t) {
                    return new Promise((i, n) => {
                        const r = new CustomEvent(e, { detail: t })
                        window.dispatchEvent(r), i()
                    })
                })
        },
        543: (e, t, i) => {
            Object.defineProperty(t, '__esModule', { value: !0 }),
                (t.getStats =
                    t.moveFile =
                    t.copyFile =
                    t.readDirectory =
                    t.removeFile =
                    t.readBinaryFile =
                    t.readFile =
                    t.writeBinaryFile =
                    t.writeFile =
                    t.removeDirectory =
                    t.createDirectory =
                        void 0)
            const n = i(69)
            ;(t.createDirectory = function (e) {
                return n.request({
                    url: 'filesystem.createDirectory',
                    type: n.RequestType.POST,
                    data: e,
                    isNativeMethod: !0,
                })
            }),
                (t.removeDirectory = function (e) {
                    return n.request({
                        url: 'filesystem.removeDirectory',
                        type: n.RequestType.POST,
                        data: e,
                        isNativeMethod: !0,
                    })
                }),
                (t.writeFile = function (e) {
                    return n.request({
                        url: 'filesystem.writeFile',
                        type: n.RequestType.POST,
                        data: e,
                        isNativeMethod: !0,
                    })
                }),
                (t.writeBinaryFile = function (e) {
                    const t = new Uint8Array(e.data)
                    let i = ''
                    for (const e of t) i += String.fromCharCode(e)
                    return n.request({
                        url: 'filesystem.writeBinaryFile',
                        type: n.RequestType.POST,
                        data: {
                            fileName: e.fileName,
                            data: window.btoa(i),
                        },
                        isNativeMethod: !0,
                    })
                }),
                (t.readFile = function (e) {
                    return n.request({
                        url: 'filesystem.readFile',
                        type: n.RequestType.POST,
                        data: e,
                        isNativeMethod: !0,
                    })
                }),
                (t.readBinaryFile = function (e) {
                    return new Promise((t, i) => {
                        n.request({
                            url: 'filesystem.readBinaryFile',
                            type: n.RequestType.POST,
                            data: e,
                            isNativeMethod: !0,
                        })
                            .then((e) => {
                                const i = window.atob(e)
                                const n = i.length
                                const r = new Uint8Array(n)
                                for (let e = 0; e < n; e++)
                                    r[e] = i.charCodeAt(e)
                                t({ data: r.buffer })
                            })
                            .catch((e) => {
                                i(e)
                            })
                    })
                }),
                (t.removeFile = function (e) {
                    return n.request({
                        url: 'filesystem.removeFile',
                        type: n.RequestType.POST,
                        data: e,
                        isNativeMethod: !0,
                    })
                }),
                (t.readDirectory = function (e) {
                    return n.request({
                        url: 'filesystem.readDirectory',
                        type: n.RequestType.POST,
                        data: e,
                        isNativeMethod: !0,
                    })
                }),
                (t.copyFile = function (e, t) {
                    return n.request({
                        url: 'filesystem.copyFile',
                        type: n.RequestType.POST,
                        data: { source: e, destination: t },
                        isNativeMethod: !0,
                    })
                }),
                (t.moveFile = function (e, t) {
                    return n.request({
                        url: 'filesystem.moveFile',
                        type: n.RequestType.POST,
                        data: { source: e, destination: t },
                        isNativeMethod: !0,
                    })
                }),
                (t.getStats = function (e) {
                    return n.request({
                        url: 'filesystem.getStats',
                        type: n.RequestType.POST,
                        data: { path: e },
                        isNativeMethod: !0,
                    })
                })
        },
        109: (e, t, i) => {
            Object.defineProperty(t, '__esModule', { value: !0 }),
                (t.init = void 0)
            const n = i(473)
            const r = i(359)
            const o = i(306)
            t.init = function () {
                if (
                    (window.NL_MODE &&
                        window.NL_MODE == 'browser' &&
                        n.ping.start(),
                    void 0 !== window.NL_ARGS)
                )
                    for (let e = 0; e < window.NL_ARGS.length; e++)
                        if (window.NL_ARGS[e] == '--debug-mode') {
                            r.devClient.start()
                            break
                        }
                window.NL_CVERSION = o.version
            }
        },
        325: (e, t, i) => {
            Object.defineProperty(t, '__esModule', { value: !0 }),
                (t.setTray =
                    t.showMessageBox =
                    t.showNotification =
                    t.showDialogSave =
                    t.showDialogOpen =
                    t.getEnvar =
                    t.execCommand =
                    t.MessageBoxType =
                        void 0)
            const n = i(69)
            let r
            ;((r = t.MessageBoxType || (t.MessageBoxType = {})).WARN = 'WARN'),
                (r.ERROR = 'ERROR'),
                (r.INFO = 'INFO'),
                (r.QUESTION = 'QUESTION'),
                (t.execCommand = function (e) {
                    return n.request({
                        url: 'os.execCommand',
                        type: n.RequestType.POST,
                        data: e,
                        isNativeMethod: !0,
                    })
                }),
                (t.getEnvar = function (e) {
                    return n.request({
                        url: 'os.getEnvar',
                        type: n.RequestType.POST,
                        data: e,
                        isNativeMethod: !0,
                    })
                }),
                (t.showDialogOpen = function (e) {
                    return n.request({
                        url: 'os.dialogOpen',
                        type: n.RequestType.POST,
                        data: e,
                        isNativeMethod: !0,
                    })
                }),
                (t.showDialogSave = function (e) {
                    return n.request({
                        url: 'os.dialogSave',
                        type: n.RequestType.POST,
                        data: e,
                        isNativeMethod: !0,
                    })
                }),
                (t.showNotification = function (e) {
                    return n.request({
                        url: 'os.showNotification',
                        type: n.RequestType.POST,
                        data: e,
                        isNativeMethod: !0,
                    })
                }),
                (t.showMessageBox = function (e) {
                    return n.request({
                        url: 'os.showMessageBox',
                        type: n.RequestType.POST,
                        data: e,
                        isNativeMethod: !0,
                    })
                }),
                (t.setTray = function (e) {
                    return n.request({
                        url: 'os.setTray',
                        type: n.RequestType.POST,
                        data: e,
                        isNativeMethod: !0,
                    })
                })
        },
        100: (e, t, i) => {
            Object.defineProperty(t, '__esModule', { value: !0 }),
                (t.getData = t.putData = void 0)
            const n = i(69)
            ;(t.putData = function (e) {
                return n.request({
                    url: 'storage.putData',
                    type: n.RequestType.POST,
                    data: e,
                    isNativeMethod: !0,
                })
            }),
                (t.getData = function (e) {
                    return n.request({
                        url: 'storage.getData',
                        type: n.RequestType.POST,
                        data: e,
                        isNativeMethod: !0,
                    })
                })
        },
        776(e, t, i) {
            const n =
                (this && this.__awaiter) ||
                function (e, t, i, n) {
                    return new (i || (i = Promise))(function (r, o) {
                        function s(e) {
                            try {
                                a(n.next(e))
                            } catch (e) {
                                o(e)
                            }
                        }
                        function u(e) {
                            try {
                                a(n.throw(e))
                            } catch (e) {
                                o(e)
                            }
                        }
                        function a(e) {
                            let t
                            e.done
                                ? r(e.value)
                                : ((t = e.value),
                                  t instanceof i
                                      ? t
                                      : new i(function (e) {
                                            e(t)
                                        })).then(s, u)
                        }
                        a((n = n.apply(e, t || [])).next())
                    })
                }
            Object.defineProperty(t, '__esModule', { value: !0 }),
                (t.create =
                    t.setSize =
                    t.setDraggableRegion =
                    t.move =
                    t.setIcon =
                    t.focus =
                    t.isVisible =
                    t.hide =
                    t.show =
                    t.isFullScreen =
                    t.exitFullScreen =
                    t.setFullScreen =
                    t.minimize =
                    t.isMaximized =
                    t.unmaximize =
                    t.maximize =
                    t.setTitle =
                        void 0)
            const r = i(69)
            ;(t.setTitle = function (e) {
                return r.request({
                    url: 'window.setTitle',
                    type: r.RequestType.POST,
                    data: { title: e },
                    isNativeMethod: !0,
                })
            }),
                (t.maximize = function () {
                    return r.request({
                        url: 'window.maximize',
                        type: r.RequestType.GET,
                        isNativeMethod: !0,
                    })
                }),
                (t.unmaximize = function () {
                    return r.request({
                        url: 'window.unmaximize',
                        type: r.RequestType.GET,
                        isNativeMethod: !0,
                    })
                }),
                (t.isMaximized = function () {
                    return r.request({
                        url: 'window.isMaximized',
                        type: r.RequestType.GET,
                        isNativeMethod: !0,
                    })
                }),
                (t.minimize = function () {
                    return r.request({
                        url: 'window.minimize',
                        type: r.RequestType.GET,
                        isNativeMethod: !0,
                    })
                }),
                (t.setFullScreen = function () {
                    return r.request({
                        url: 'window.setFullScreen',
                        type: r.RequestType.GET,
                        isNativeMethod: !0,
                    })
                }),
                (t.exitFullScreen = function () {
                    return r.request({
                        url: 'window.exitFullScreen',
                        type: r.RequestType.GET,
                        isNativeMethod: !0,
                    })
                }),
                (t.isFullScreen = function () {
                    return r.request({
                        url: 'window.isFullScreen',
                        type: r.RequestType.GET,
                        isNativeMethod: !0,
                    })
                }),
                (t.show = function () {
                    return r.request({
                        url: 'window.show',
                        type: r.RequestType.GET,
                        isNativeMethod: !0,
                    })
                }),
                (t.hide = function () {
                    return r.request({
                        url: 'window.hide',
                        type: r.RequestType.GET,
                        isNativeMethod: !0,
                    })
                }),
                (t.isVisible = function () {
                    return r.request({
                        url: 'window.isVisible',
                        type: r.RequestType.GET,
                        isNativeMethod: !0,
                    })
                }),
                (t.focus = function () {
                    return r.request({
                        url: 'window.focus',
                        type: r.RequestType.GET,
                        isNativeMethod: !0,
                    })
                }),
                (t.setIcon = function (e) {
                    return r.request({
                        url: 'window.setIcon',
                        type: r.RequestType.POST,
                        isNativeMethod: !0,
                        data: { icon: e },
                    })
                }),
                (t.move = function (e, t) {
                    return r.request({
                        url: 'window.move',
                        type: r.RequestType.POST,
                        isNativeMethod: !0,
                        data: { x: e, y: t },
                    })
                }),
                (t.setDraggableRegion = function (e) {
                    return new Promise((t, i) => {
                        const r = document.getElementById(e)
                        let o = 0
                        let s = 0
                        function u(e) {
                            return n(this, void 0, void 0, function* () {
                                yield Neutralino.window.move(
                                    e.screenX - o,
                                    e.screenY - s,
                                )
                            })
                        }
                        r || i(`Unable to find dom element: #${e}`),
                            r.addEventListener('mousedown', (e) => {
                                ;(o = e.clientX),
                                    (s = e.clientY),
                                    r.addEventListener('mousemove', u)
                            }),
                            r.addEventListener('mouseup', () => {
                                r.removeEventListener('mousemove', u)
                            }),
                            t()
                    })
                }),
                (t.setSize = function (e) {
                    return r.request({
                        url: 'window.setSize',
                        type: r.RequestType.POST,
                        isNativeMethod: !0,
                        data: e,
                    })
                }),
                (t.create = function (e, t) {
                    return new Promise((i, n) => {
                        function r(e) {
                            return (
                                typeof e !== 'string' ||
                                    ((e = e.trim()).includes(' ') &&
                                        (e = `"${e}"`)),
                                e
                            )
                        }
                        let o = window.NL_ARGS.reduce(
                            (e, t, i) => (
                                (t.includes('--path=') ||
                                    t.includes('--debug-mode') ||
                                    t.includes('--load-dir-res') ||
                                    i == 0) &&
                                    (e += ` ${r(t)}`),
                                e
                            ),
                            '',
                        )
                        o += ` --url=${r(e)}`
                        for (const e in t)
                            e != 'processArgs' &&
                                (o += ` --window${e.replace(
                                    /[A-Z]|^[a-z]/g,
                                    (e) => `-${e.toLowerCase()}`,
                                )}=${r(t[e])}`)
                        t.processArgs && (o += ` ${t.processArgs}`),
                            Neutralino.os
                                .execCommand({
                                    command: o,
                                    shouldRunInBackground: !0,
                                })
                                .then(() => {
                                    i()
                                })
                                .catch((e) => {
                                    n(e)
                                })
                    })
                })
        },
        359(e, t, i) {
            const n =
                (this && this.__awaiter) ||
                function (e, t, i, n) {
                    return new (i || (i = Promise))(function (r, o) {
                        function s(e) {
                            try {
                                a(n.next(e))
                            } catch (e) {
                                o(e)
                            }
                        }
                        function u(e) {
                            try {
                                a(n.throw(e))
                            } catch (e) {
                                o(e)
                            }
                        }
                        function a(e) {
                            let t
                            e.done
                                ? r(e.value)
                                : ((t = e.value),
                                  t instanceof i
                                      ? t
                                      : new i(function (e) {
                                            e(t)
                                        })).then(s, u)
                        }
                        a((n = n.apply(e, t || [])).next())
                    })
                }
            Object.defineProperty(t, '__esModule', { value: !0 }),
                (t.devClient = void 0)
            const r = i(69)
            t.devClient = {
                start() {
                    setInterval(
                        () =>
                            n(this, void 0, void 0, function* () {
                                try {
                                    ;(yield r.request({
                                        url: 'http://localhost:5050',
                                        type: r.RequestType.GET,
                                    })).needsReload && location.reload()
                                } catch (e) {
                                    console.error(
                                        'Unable to communicate with neu devServer',
                                    )
                                }
                            }),
                        1e3,
                    )
                },
            }
        },
        69: (e, t) => {
            let i
            Object.defineProperty(t, '__esModule', { value: !0 }),
                (t.request = t.RequestType = void 0),
                ((i = t.RequestType || (t.RequestType = {})).GET = 'GET'),
                (i.POST = 'POST'),
                (t.request = function (e) {
                    return new Promise((t, i) => {
                        let n = null
                        const r = (function () {
                            let e
                            return (
                                (e = window.XMLHttpRequest
                                    ? new XMLHttpRequest()
                                    : new ActiveXObject('Microsoft.XMLHTTP')),
                                e
                            )
                        })()
                        ;(r.onreadystatechange = () => {
                            if (r.readyState == 4 && r.status == 200) {
                                let e = null
                                const n = r.responseText
                                n && (e = JSON.parse(n)),
                                    e &&
                                        e.success &&
                                        (e.hasOwnProperty('returnValue')
                                            ? t(e.returnValue)
                                            : t(e)),
                                    e && e.error && i(e.error)
                            } else
                                r.readyState == 4 &&
                                    i(
                                        'Neutralino server is offline. Try restarting the application',
                                    )
                        }),
                            e.isNativeMethod &&
                                (e.url = `http://localhost:${window.NL_PORT}/__nativeMethod_${e.url}`),
                            e.data && (n = JSON.stringify(e.data)),
                            r.open(e.type, e.url, !0),
                            r.setRequestHeader(
                                'Content-Type',
                                'application/json',
                            ),
                            r.setRequestHeader(
                                'Authorization',
                                `Basic ${window.NL_TOKEN}`,
                            ),
                            r.send(n)
                    })
                })
        },
        473(e, t, i) {
            const n =
                (this && this.__awaiter) ||
                function (e, t, i, n) {
                    return new (i || (i = Promise))(function (r, o) {
                        function s(e) {
                            try {
                                a(n.next(e))
                            } catch (e) {
                                o(e)
                            }
                        }
                        function u(e) {
                            try {
                                a(n.throw(e))
                            } catch (e) {
                                o(e)
                            }
                        }
                        function a(e) {
                            let t
                            e.done
                                ? r(e.value)
                                : ((t = e.value),
                                  t instanceof i
                                      ? t
                                      : new i(function (e) {
                                            e(t)
                                        })).then(s, u)
                        }
                        a((n = n.apply(e, t || [])).next())
                    })
                }
            Object.defineProperty(t, '__esModule', { value: !0 }),
                (t.ping = void 0)
            const r = i(885)
            t.ping = {
                start: () => {
                    setInterval(
                        () =>
                            n(void 0, void 0, void 0, function* () {
                                yield r.keepAlive()
                            }),
                        5e3,
                    )
                },
            }
        },
        306: (e) => {
            e.exports = JSON.parse(
                '{"name":"neutralino-client-library","version":"1.5.0","description":"","main":"index.js","scripts":{"test":"echo \\"Error: no test specified\\" && exit 1","build":"webpack","watch":"webpack --watch"},"repository":{"type":"git","url":"git+https://github.com/neutralinojs/neutralino.js.git"},"author":"Neutralinojs","license":"MIT","bugs":{"url":"https://github.com/neutralinojs/neutralino.js/issues"},"homepage":"https://github.com/neutralinojs/neutralino.js#readme","devDependencies":{"@types/node":"^16.4.2","ts-loader":"^9.2.4","typescript":"^4.3.5","webpack":"^5.46.0","webpack-cli":"^4.7.2"},"dependencies":{}}',
            )
        },
    }
    const t = {}
    function i(n) {
        const r = t[n]
        if (void 0 !== r) return r.exports
        const o = (t[n] = { exports: {} })
        return e[n].call(o.exports, o, o.exports, i), o.exports
    }
    const n = {}
    ;(() => {
        const e = n
        Object.defineProperty(e, '__esModule', { value: !0 }),
            (e.init =
                e.events =
                e.window =
                e.app =
                e.debug =
                e.storage =
                e.computer =
                e.os =
                e.filesystem =
                    void 0),
            (e.filesystem = i(543)),
            (e.os = i(325)),
            (e.computer = i(308)),
            (e.storage = i(100)),
            (e.debug = i(199)),
            (e.app = i(885)),
            (e.window = i(776)),
            (e.events = i(284))
        const t = i(109)
        Object.defineProperty(e, 'init', {
            enumerable: !0,
            get() {
                return t.init
            },
        })
    })(),
        (Neutralino = n)
})()
/* eslint-enable */
